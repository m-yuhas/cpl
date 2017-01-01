// Main Function and Entry Point into CPL
// (C) 2016 Michael Yuhas

#include <iostream>
#include <cstdlib>
#include <fstream>
#include <string>
#include <vector>
#include "../include/StringUtils.hpp"

int main( int argc, char *argv[] )
{
  // Check Number Of Arguments, Make sure a filename is present
  if ( argc < 2 )
  {
    std::cout << "使用法：$cpl <编程的名称>\n";
    return EXIT_SUCCESS;
  }
  // Open File and Read Lines into an array of Strings
  std::string fileName = argv[1];
  std::ifstream ifstr;
  ifstr.open( fileName.c_str() );
  std::string line;
  std::vector<std::string> lineVector;
  while ( std::getline( ifstr, line) )
  {
    lineVector.push_back(line);
  }
  ifstr.close();


  for ( std::vector<std::string>::iterator it = lineVector.begin(); it != lineVector.end(); it++ ) {
    *it = StringUtils::stripWhiteSpace( *it );
    std::cout << *it;
  }


  return EXIT_SUCCESS;
}
