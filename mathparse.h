#include <stdlib.h>
#include "hashmap.h"

#define IDLE_STATE 0
#define PARENTH_STATE 1

int evaluate_expression( char *expression );

int evaluate__int_expression( char *expression ) {
  //First evaluate parenthesis
  char state = IDLE_STATE;
  for ( int i = 0; i < strlen(expression); i++ ) {
    if ( expression[i] == '(' ) {
      state=PARENTH_STATE;
    }

  }

  //Then Multiplication and division

  //Then Addition and Subtraction
}
