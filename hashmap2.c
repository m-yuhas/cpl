// stack.h
// This library implements a stack in C.  The stack is a stack of integers.

#include <stdlib.h>
#include <string.h>

#define MAP_SIZE 1024

struct HashNode {
  const char *key;
  const char type;
  const char *pointerToString
  float Float
  int Int;
};

struct Hashmap {
  struct HashNode HashNodeArray[MAP_SIZE];
}

int add_key_value_pair( struct Hashmap *pointerToHashmap, const char *key, struct Var value );
int get_value_at_key( struct Hashmap *pointerToHashmap, const char *key, struct Var *value );
int destroy_hashmap( struct Hashmap *pointerToHashmap );
int hash( const char *key, unsigned int *hashvalue );

int add_key_value_pair_int( struct Hashmap *pointerToHashmap, const char *key, int value ) {
  // This function stores a key value pair
  // Returns 0 for success
  // Returns 1 if the map is full
  unsigned int hashValue;
  hash( key, &hashValue );
  hashValue = hashValue % MAP_SIZE;
  if ( HashNodeArray[hashValue].type == 0 ) {
    HashNodeArray[hashValue].type = 1;
    HashNodeArray[hashValue].Int = value;
    HashNodeArray[hashValue].key = *key;
    return 0;
  } else {
    for ( int i = hashValue + 1; i < MAP_SIZE; i++ ) {
      if ( HashNodeArray[i].type == 0 ) {
        HashNodeArray[i].type = 1;
        HashNodeArray[i].Int = value;
        HashNodeArray[i].key = *key;
        return 0;
      }
    }
    for ( int i = 0; i < hashValue; i++ ) {
      if ( HashNodeArray[i].type == 0 ) {
        HashNodeArray[i].type = 1;
        HashNodeArray[i].Int = value;
        HashNodeArray[i].key = *key;
        return 0;
      }
    }
    return 1;
  }
}

int get_key_type( struct Hashmap *pointerToHashmap, const char *key ) {
  unsigned int hashValue;
  hash( key, &hashValue );
  hashValue = hashValue % MAP_SIZE;
  if ( strcmp(HashNodeArray[hashValue].key,*key) ) {
    return HashNodeArray[hashValue].type;
  } else {
    for ( int i = hashValue + 1; i < MAP_SIZE; i++ ) {
      if ( strcmp(HashNodeArray[i].key,*key) ) {
        return HashNodeArray[i].type;
      }
    }
    for ( int i = 0; i < hashValue; i++ ) {
      if ( strcmp(HashNodeArray[i].key,*key) ) {
        return HashNodeArray[i].type;
      }
    }
    return -1;
  }
}

int get_int_at_key( struct Hashmap *pointerToHashmap, const char *key, int *value ) {
  unsigned int hashValue;
  hash( key, &hashValue );
  hashValue = hashValue % MAP_SIZE;
  if ( strcmp(HashNodeArray[hashValue].key,*key) ) {
    *value = HashNodeArray[hashValue];
    return 0;
  } else {
    for ( int i = hashValue + 1; i < MAP_SIZE; i++ ) {
      if ( strcmp(HashNodeArray[i].key,*key) ) {
        *value = HashNodeArray[i];
        return 0;
      }
    }
    for ( int i = 0; i < hashValue; i++ ) {
      if ( strcmp(HashNodeArray[i].key,*key) ) {
        *value = HashNodeArray[i];
        return 0;
      }
    }
    return 1;
  }
}

int hash( const char *key, unsigned int *hashValue) {
  // This function hashes some key.
  // The hash is written to the integer: hashValue
  // Returns 0 for success
  *hashValue = 0x55555555;
  while ( strcmp(key,"\0") != 0 ) {
    *hashValue = ( ( *hashValue << 5 ) | ( ( *hashValue & 0xF8000000 ) >> 27 ) ^ *key++ );
  }
  return 0;
}
