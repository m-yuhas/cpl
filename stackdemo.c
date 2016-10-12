#include <stdio.h>
#include "stack.h"

int main() {
  struct Stack stack1;

  printf("Created stack var\n");
  init_stack( &stack1 );
  printf("stack initialized\n");
  push( &stack1, 10 );
  printf("stack pushed\n");
  int a;
  pop( &stack1, &a );
  printf("stack popped\n");
  printf("a's value is %d\n",a);

  struct Stack stack2;
  init_stack( &stack2 );
  for ( int i=12; i < 20; i++ ) {
    printf("%d\n",stack2.currentLength);
    push(&stack2, i);
  }
  int b;
  for ( int i=12; i < 20; i++ ) {
    pop(&stack2, &b);
    printf("popped %d\n",b);
  }

  destroy_stack(&stack1);
  destroy_stack(&stack2);

  return 0;
}
