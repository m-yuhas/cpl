#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <ctype.h>
#include "hashmap.h"

//TODO Remove Whitspace From expression
//TODO Handle Parenthesis

int evaluate_int_expression( char *expression, struct Hashmap *pointerToHashMap, int *value );
int eval_atom( char *expression, struct Hashmap *pointerToHashMap, int *value );

int evaluate_int_expression( char *expression, struct Hashmap *pointerToHashMap, int *value ) {
  char addsubindex = -1;
  char muldivindex = -1;
  char optype = -1;
  for ( int i = 0; i < strlen(expression); i++ ) {
    if ( expression[i] == '+' || expression[i] == '-' ) {
      addsubindex = i;
      if ( expression[i] == '+' ) {
        optype = 1;
      } else {
        optype = 2;
      }
      break;
    }
    if ( expression[i] == '*' || expression[i] == '/' ) {
      muldivindex = i;
      if ( expression[i] == '*' ) {
        optype = 1;
      } else {
        optype = 2;
      }
      break;
    }
  }
  if ( addsubindex != -1 ) {
    char firsthalf[strlen(expression)];
    char lasthalf[strlen(expression)];
    strncpy(firsthalf,expression,addsubindex);
    strncpy(lasthalf,expression+addsubindex+1,strlen(expression)-addsubindex+1);
    int returnVal1;
    int returnVal2;
    int error;
    error = evaluate_int_expression( firsthalf, pointerToHashMap, &returnVal1 );
    error = evaluate_int_expression( lasthalf, pointerToHashMap, &returnVal2 );
    if ( optype == 1 ) {
      *value = returnVal1 + returnVal2;
      return error;
    } else {
      *value = returnVal1 - returnVal2;
      return error;
    }
  }
  if ( muldivindex != -1 ) {
    printf("Multiply Divide Reached\n");
    char firsthalf[strlen(expression)];
    char lasthalf[strlen(expression)];
    strncpy(firsthalf,expression,muldivindex);
    strncpy(lasthalf,expression+muldivindex+1,strlen(expression)-muldivindex+1);
    int returnVal1;
    int returnVal2;
    int error;
    error = evaluate_int_expression( firsthalf, pointerToHashMap, &returnVal1 );
    error = evaluate_int_expression( lasthalf, pointerToHashMap, &returnVal2 );
    if ( optype == 1 ) {
      *value = returnVal1 * returnVal2;
      return error;
    } else {
      *value = returnVal1 / returnVal2;
      return error;
    }
  }
  return eval_atom( expression, pointerToHashMap, value );
}

int eval_atom( char *expression, struct Hashmap *pointerToHashMap, int *value ) {
  if ( isdigit(expression[0]) ) {
    *value = atoi(expression);
    return 0;
  } else {
    return get_int_at_key(pointerToHashMap,expression,value);
  }
}
