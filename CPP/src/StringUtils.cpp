// String Utilities
// (C) 2016 Michael Yuhas

#include <string>
#include "../include/StringUtils.hpp"

std::string StringUtils::stripWhiteSpace( std::string inputString )
{
  int startIndex = 0;
  for ( std::string::iterator it = inputString.begin(); it != inputString.end(); it++ )
  {
    if ( *it == '\t' || *it == ' ' )
    {
      startIndex++;
    }
    else
    {
      break;
    }
  }
  int endIndex = inputString.length()-1;
  for ( std::string::iterator it = inputString.end(); it != inputString.begin(); it-- )
  {
    if ( *it == '\t' || *it == ' ' )
    {
      endIndex--;
    }
    else
    {
      break;
    }
  }
  if ( startIndex >= endIndex ) {
    return "";
  }
  return inputString.substr( startIndex, endIndex );
}

std::string StringUtils::removeWhiteSpaceNotInQuotes( std::string inputString )
{
  return inputString;
}
