// String Utilities
// (C) 2016 Michael Yuhas

#include <string>
#include <sstream>
#include "../include/StringUtils.hpp"

/*
stripWhiteSpace: Removes whitespace from beginning and end of string
Arguments:
inputString - String from which to remove whitespace
Returns: A string with whitespace removed
*/
std::string StringUtils::stripWhiteSpace( std::string inputString )
{
  unsigned int startIndex = 0;
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
  unsigned int charsFromEnd = 0;
  for ( std::string::iterator it = inputString.end(); it != inputString.begin(); it-- )
  {
    // Check if character is whitespace
    if ( *it == '\t' || *it == ' ' )
    {
      charsFromEnd++;
    }
    else
    {
      break;
    }
  }
  // If the string is all whitespace return an empty string
  if ( startIndex >= inputString.length()-2-charsFromEnd ) {
    return "";
  }
  return inputString.substr( startIndex, inputString.length()-2-charsFromEnd-startIndex );
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

std::stod StringUtils::removeWhiteSpaceNotInQuotes( std::string inputString )
{
  double returnDouble = 0;
  for ( string::iterator it = inputString.begin(); it < inputString.end(); it++ )
  {
    if ( *it == '.' ) //TODO Include Chinese Period
    {
      break;
    }
  }
  string::iterator it_reverse = it;
  string::iterator it_forward = it;
  int place = 1;
  for ( it_reverse; it_reverse >= inputString.end(); it_revers-- )
  {
    switch (*it)
    {
      case '0':
      {
        break;
      }
      case '1':
      {
        returnDouble += 1*place;
        break;
      }
      case '2':
      {
        returnDouble += 2*place;
        break;
      }
      case '3':
      {
        returnDouble += 3*place;
        break;
      }
      case '4':
      {
        returnDouble += 4*place;
        break;
      }
      case '5':
      {
        returnDouble += 5*place;
        break;
      }
      case '6':
      {
        returnDouble += 6*place;
        break;
      }
      case '7':
      {
        returnDouble += 7*place;
        break;
      }
      case '8':
      {
        returnDouble += 8*place;
        break;
      }
      case '9':
      {
        returnDouble += 9*place;
        break;
      }
      default:
      {
        throw //TODO create exception
      }
    }
    place *= 10;
  }
  place = 0.1;
  

}
