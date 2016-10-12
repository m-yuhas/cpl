// stack.h
// This library implements a stack in C.  The stack is a stack of integers.
// TODO: Make this library extensible for other data types and structs

#include <stdlib.h>
#include <string.h>

struct Var {
  const char type;
  const char *pointerToString;
  float *pointerToFloat;
  long *pointerToInt;
};

struct HashNode {
  struct  Var *pointerToStoredVar;
  const char *key;
};

struct Hashmap {
  struct HashNode *pointerToNodeArray;
  int size;
};

int init_hashmap( struct Hashmap *pointerToHashmap, int size );
int add_key_value_pair( struct Hashmap *pointerToHashmap, const char *key, struct Var value );
int get_value_at_key( struct Hashmap *pointerToHashmap, const char *key, struct Var *value );
int destroy_hashmap( struct Hashmap *pointerToHashmap );
int hash( const char *key, unsigned int *hashvalue );

int init_hashmap( struct Hashmap *pointerToHashmap, int size ) {
  pointerToHashmap->pointerToNodeArray = malloc( size*sizeof(struct HashNode) );
  pointerToHashmap->size = size;
  return 0;
}

int add_key_value_pair( struct Hashmap *pointerToHashmap, const char *key, struct Var value ) {
  // This function stores a key value pair
  // Returns 0 for success
  // Returns 1 if the map is full
  unsigned int hashValue;
  hash( key, &hashValue );
  hashValue = hashValue % pointerToHashmap->size;
  if ( *(pointerToHashmap->pointerToNodeArray + sizeof(hashValue)*hashValue)->pointerToStoredVar->type == 0 ) {
    *(pointerToHashmap->pointerToNodeArray+hashValue)->pointerToStoredVar = value;
    *(pointerToHashmap->pointerToNodeArray+hashValue)->key = *key;
    return 0;
  } else {
    for ( int i = hashValue + 1; i < size; i++ ) {
      if ( *(pointerToHashmap->pointerToNodeArray + i) == NULL ) {
        *(pointerToHashmap->pointerToNodeArray+i)->pointerToStoredVar = value;
        *(pointerToHashmap->pointerToNodeArray+i)->key = *key;
        return 0;
      }
    }
    for ( int i = 0; i < hashValue; i++ ) {
      if ( *(pointerToHashmap->pointerToNodeArray + i) == NULL ) {
        *(pointerToHashmap->pointerToNodeArray+i)->pointerToStoredVar = value;
        *(pointerToHashmap->pointerToNodeArray+i)->key = *key;
        return 0;
      }
    }
    return 1;
  }
}

int get_value_at_key( struct Hashmap *pointerToHashmap, const char *key, struct Var *value ) {
  // This function retreives the value at a given key
  // Returns 0 for success
  // Returns 1 if the key is not found in the map
  unsigned int hashValue;
  hash( key, &hashValue );
  hashValue = hashValue % size;
  if ( *(pointerToHashmap->pointerToNodeArray + hashValue)->key == *key ) {
    *value = *(pointerToHashmap->pointerToNodeArray+hashValue)->pointerToStoredVar;
    return 0;
  } else {
    for ( int i = hashValue + 1; i < size; i++ ) {
      if ( *(pointerToHashmap->pointerToNodeArray + i)->key == *key ) {
        *value = *(pointerToHashmap->pointerToNodeArray+i)->pointerToStoredVar;
        return 0;
      }
    }
    for ( int i = 0; i < hashValue; i++ ) {
      if ( *(pointerToHashmap->pointerToNodeArray + i)->key == *key ) {
        *value = *(pointerToHashmap->pointerToNodeArray+i)->pointerToStoredVar;
        return 0;
      }
    }
    return 1;
  }
}

int destroy_hashmap( struct Hashmap *pointerToHashmap ) {
  // Clean up when destroying Hashmap
  free( pointerToHashmap->pointerToNodeArray );
  return 0;
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
