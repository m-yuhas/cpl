#include <iostream>
#include <fstream>
#include <cstring>
#include <string>
#include "utils.hpp"

using namespace std;

int main( int argc, char* argv[] ) {
  //first parse command line arguments
  if ( argc != 2 ) {
    cout << "没有程序：请包括程序\n";
    return 0;
  }

  int numberOfLines = 0;
  int maxLineSize = 0;
  string line;
  ifstream scriptFile( argv[1] );
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
    string thisLine = Utils::stripWhiteSpace(programBuffer[i]);
    if ( thisLine.length() == 0 ) {
      continue;
    }
    if ( thisLine.find("备注:") == 0 ) {
      continue;
    }
    if ( thisLine.find("印\"") == 0 ) {
      string printString = Utils::print(thisLine.substr(3));
      if ( printString.find("错误") == 0 ) {
        break;
      } else {
        cout << printString << "\n";
        continue;
      }
    }
    cout << "\033[1;31m错误：不清楚指令在" << i << "号句\033[0m\n" << programBuffer[i] << "\n";
    break;
    //cout << Utils::stripWhiteSpace(programBuffer[i]) << "\n";
    //cout << programBuffer[i] << "\n";
  }

  return 0;

}
