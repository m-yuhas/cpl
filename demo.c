#include <stdio.h>
#include <string.h>

int main( int argc, char *argv[] ) {
  //Parse Input Arguments
  if ( argc < 2 ) {
    printf("Please Include A Script To Execute...\n");
    return -1;
  }

  //Open File and Read Contents into a Buffer
  FILE *filePointer;
  filePointer = fopen(argv[1],"r");
  if ( filePointer == NULL ) {
    printf("Error Opening File\n");
    return -1;
  }
  int c;
  int charCount = 0;
  int lineCount = 0;
  while ( c != EOF ) {
    c = fgetc( filePointer );
    charCount++;
    if ( c == '\n' ) {
      lineCount++;
    }
    if ( c == '创' ) {
      printf("Found unicode Character");
    }
    //TODO: Check for overflow here
  }
  fseek( filePointer, SEEK_SET, 0 );
  char fileBuffer[charCount];
  fread(fileBuffer, sizeof(fileBuffer), 1, filePointer);
  fclose( filePointer );

  //Find the location of NewLines in The File
  int lineBreakArray[lineCount];
  int currLine = 1;
  for ( int i=0; i < charCount; i++ ) {
    if ( fileBuffer[i] == '\n' ) {
      lineBreakArray[currLine] = i;
      currLine++;
    }
  }

  //Parse The Script
  for ( int line=0; line <= lineCount; line++ ) {
    int endOfLinePos;
    if ( line == lineCount ) {
      endOfLinePos = charCount;
    } else {
      endOfLinePos = lineBreakArray[line+1];
    }
    for ( int cursor=lineBreakArray[line]; cursor < endOfLinePos; cursor++ ) {
      if ( fileBuffer[cursor] == '创' ) {
        printf("Created a variable here\n");
      }
    }
  }


  // Return with No Error
  return 0;
}
