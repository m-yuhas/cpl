#include "mathparse.h"
#include <stdio.h>

int main() {
  struct Hashmap hashmap1;
  init_hashmap( &hashmap1 );

  add_key_value_pair_int( &hashmap1, "地土", 2 );

  int value;

  int result = evaluate_int_expression("4+4",&hashmap1,&value);

  printf("Result is:%d\n",value);

  result = evaluate_int_expression("2-2",&hashmap1,&value);
  printf("Result is:%d\n",value);

  result = evaluate_int_expression("2*3",&hashmap1,&value);
  printf("Result is:%d\n",value);

  result = evaluate_int_expression("8/4",&hashmap1,&value);
  printf("Result is:%d\n",value);

  result = evaluate_int_expression("地土+2",&hashmap1,&value);
  printf("Result is %d\n",value);

  result = evaluate_int_expression("地土+2-3",&hashmap1,&value);
  printf("Result is %d\n",value);

  result = evaluate_int_expression("地土+2*3",&hashmap1,&value);
  printf("Result is %d\n",value);
  return 0;
}
