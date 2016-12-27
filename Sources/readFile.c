// This is a c function that reads a file at the input pathToFile
// (C) 2016 Michael Yuhas
#include <stdio.h>
#include <stdlib.h>

char* readFile( const char* pathToFile ) {
  FILE *filePointer;
  filePointer = fopen( pathToFile, "r");
  if ( filePointer == NULL ) {
    //open file error
    return "ERR001";
  }
  char* fileBuffer;
  fileBuffer = (char*) malloc(0);
  int count = 0;
  int temp;
  while ( 1 ) {
    temp = fgetc( filePointer );
    if ( temp == EOF ) {
      break;
    }
    fileBuffer = realloc(fileBuffer, sizeof(char)*(count+1));
    if ( fileBuffer == NULL ) {
      //memory allocation Error
      free( fileBuffer );
      return "ERR002";
    }
    *(fileBuffer+count) = (char) temp;
    count++;
  }
  fclose( filePointer );
  return fileBuffer;
}
