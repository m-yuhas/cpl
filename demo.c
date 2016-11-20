#include <stdio.h>
#include <string.h>
#include "stack.h"
#include "mathparse.h"

int main( int argc, char *argv[] ) {
  //Parse Input Arguments
  if ( argc < 2 ) {
    printf("请包括程序的文件名称...\n");
    return -1;
  }

  //Open File and Read Contents into a Buffer
  FILE *filePointer;
  filePointer = fopen(argv[1],"r");
  if ( filePointer == NULL ) {
    printf("错误开文件\n");
    return -1;
  }
  int c;
  int charCount = 0;
  int lineCount = 0;
  int whiteSpaceCount = 0;
  while ( c != EOF ) {
    c = fgetc( filePointer );
    charCount++;
    if ( c == '\n' ) {
      lineCount++;
    }
    if ( c == ' ' || c == '\t' ) {
      whiteSpaceCount++;
    }
    //TODO: Check for overflow here
  }
  fseek( filePointer, SEEK_SET, 0 );
  char fileBuffer[charCount];
  char programBuffer[charCount-whiteSpaceCount];
  fread(fileBuffer, sizeof(fileBuffer), 1, filePointer);
  fclose( filePointer );

  //Copy the File sans whitespace to programBuffer
  int lineBreakArray[lineCount];
  int currLine = 1;
  int fcount=0;
  int pcount=0;
  while ( fcount < charCount ) {
    if ( fileBuffer[fcount] != ' ' && fileBuffer[fcount] != '\t' ) {
      programBuffer[pcount] = fileBuffer[fcount];
      pcount++;
    }
    if ( fileBuffer[fcount] == '\n' ) {
      lineBreakArray[currLine] = pcount;
      currLine++;
    }
    fcount++;
  }
  //Parse The Script
  for ( int line=0; line <= lineCount; line++ ) {
    printf("here\n");
    int endOfLinePos;
    if ( line == lineCount ) {
      endOfLinePos = charCount-whiteSpaceCount;
    } else {
      endOfLinePos = lineBreakArray[line+1];
      printf("not last line\n");
    }
    printf("%d\n",endOfLinePos);
    printf("%d\n",lineBreakArray[line+1]);
    char lineString[endOfLinePos - lineBreakArray[line+1]];
    printf("1\n");
    strncpy(lineString,programBuffer+lineBreakArray[line],endOfLinePos - lineBreakArray[line]);
    //for ( int cursor=lineBreakArray[line]; cursor < endOfLinePos; cursor++ ) {
      //printf("%c",programBuffer[cursor]);
    //}
    printf("2\n");
    if (programBuffer[lineBreakArray[line]] == '\n') {
      continue;
    }
    if ( memcmp(lineString,"印",2) == 0 ) {
      printf("here\n");
      printf("%s",lineString+2);
    }
  }


  // Return with No Error
  return 0;
}
