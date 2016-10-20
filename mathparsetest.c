#include "mathparse.h"
#include <stdio.h>

int main() {
  struct Hashmap hashmap1;
  init_hashmap( &hashmap1 );

  printf("HERE\n");

  add_key_value_pair_int( &hashmap1, "地土", 2 );

  int value;

  printf("Here 2\n");
  int result = evaluate_int_expression("4+4",&hashmap1,&value);

  printf("Result is:%d\n",value);
  return 0;
}
