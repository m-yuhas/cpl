#include <iostream>
#include <fstream>
#include <cstring>
#include <string>

using namespace std;

int main( int argc, char* argv[] ) {
  //first parse command line arguments
  if ( argc != 1 ) {
    cout << "没有程序：请包括程序\n";
    return 0;
  }

  int numberOfLines = 0;
  int maxLineSize = 0;
  string line;
  ifstream scriptFile( argv[0] );
  while ( getline(scriptFile, line) ) {
    numberOfLines++;
    if ( line.length() > maxLineSize ) {
      maxLineSize = line.length();
    }
  }


  //char programBuffer[numberOfLines][maxLineSize];
  string programBuffer[numberOfLines];
  scriptFile.clear();
  scriptFile.seekg( 0, ios::beg );

  int count = 0;
  while ( getline(scriptFile, line) ) {
    //strncpy( programBuffer[count], line.c_str(), line.length() );
    programBuffer[count].assign(line);
    count++;
  }
  scriptFile.close();

  for ( int i = 0; i < numberOfLines; i++ ) {
    cout << programBuffer[i] << "\n";
  }

  return 0;

}
