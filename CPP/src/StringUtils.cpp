// String Utilities
// (C) 2016 Michael Yuhas

#include <string>
#include "../include/StringUtils.hpp"

/*
stripWhiteSpace: Removes whitespace from beginning and end of string
Arguments:
inputString - String from which to remove whitespace
Returns: A string with whitespace removed
*/
std::string StringUtils::stripWhiteSpace( std::string inputString )
{
  int startIndex = 0;
  // Iterate from front of string until non whitespace character found
  for ( std::string::iterator it = inputString.begin(); it != inputString.end(); it++ )
  {
    // Check if character is whitespace
    if ( *it == '\t' || *it == ' ' )
    {
      startIndex++;
    }
    else
    {
      break;
    }
  }
  // Iterate from back of string until non whitespace character found
  int endIndex = inputString.length()-1;
  for ( std::string::iterator it = inputString.end(); it != inputString.begin(); it-- )
  {
    // Check if character is whitespace
    if ( *it == '\t' || *it == ' ' )
    {
      endIndex--;
    }
    else
    {
      break;
    }
  }
  // If the string is all whitespace return an empty string
  if ( startIndex >= endIndex ) {
    return "";
  }
  return inputString.substr( startIndex, endIndex );
}

/*
removeWhiteSpaceNotInQuotes: Remove whitespace not contained in quotation marks
Arguments:
inputString - String from which to remove whitespace
Returns: String with whitespace removed
*/
std::string StringUtils::removeWhiteSpaceNotInQuotes( std::string inputString )
{
  return inputString;
}
