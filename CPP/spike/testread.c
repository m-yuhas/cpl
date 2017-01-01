#include <stdio.h>
#include "readFile.c"

int main() {
  char * fname = "test.txt";
  char * output = readFile(fname);
  printf("%s",output);
  return 0;
}
