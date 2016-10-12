// stack.h
// This library implements a stack in C.  The stack is a stack of integers.
// TODO: Make this library extensible for other data types and structs

#include <stdlib.h>

struct Stack {
  int *pointerToArray;
  int currentLength;
};

int init_stack( struct Stack *pointerToStack );
int push( struct Stack *pointerToStack, int value );
int pop( struct Stack *pointerToStack, int *value );
int destroy_stack( struct Stack *pointerToStack );

int init_stack( struct Stack *pointerToStack ) {
  // Initializes a Stack with pointerToStack as the pointer to that Stack
  // Returning anything other than zero is considered a failure
  pointerToStack->pointerToArray = malloc( 0 );
  if ( pointerToStack->pointerToArray == NULL ) {
    return 1;
  }
  pointerToStack->currentLength = 0;
  return 0;
}

int push( struct Stack *pointerToStack, int value ) {
  // This pushes a value on the stack at &pointerToStack
  // Returning 0 indicates a success
  pointerToStack->currentLength++;
  pointerToStack->pointerToArray = realloc( pointerToStack->pointerToArray, pointerToStack->currentLength*sizeof(int) );
  if ( pointerToStack->pointerToArray == NULL ) {
    return 1;
  }
  *(pointerToStack->pointerToArray+pointerToStack->currentLength-1) = value;
  return 0;
}

int pop( struct Stack *pointerToStack, int *value ) {
  // This pops a value from the stack at &pointerToStack
  // Returning 0 indicates a success
  // Returning 1 idicates a memory allocation Error
  // Returning 2 indicates that the stack as no elements on it
  if ( pointerToStack->currentLength == 0 ) {
    return 2;
  }
  pointerToStack->currentLength--;
  *value = *(pointerToStack->pointerToArray+pointerToStack->currentLength);
  pointerToStack->pointerToArray = realloc( pointerToStack->pointerToArray, pointerToStack->currentLength*sizeof(int) );
  if ( pointerToStack->pointerToArray == NULL ) {
    return 1;
  }
  return 0;
}

int destroy_stack( struct Stack *pointerToStack ) {
  // This function cleans up after the stack is done being used
  free( pointerToStack->pointerToArray );
  return 0;
}
