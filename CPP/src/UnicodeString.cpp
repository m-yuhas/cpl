// Test Class to Implement Unicode Strings in CPP
// (C) 2016 Michael Yuhas

#include <vector>
#include <string>
#include <iostream>
#include "../include/UnicodeString.hpp"
#include "../include/InvalidUTF8Exception.hpp"

/*
UnicodeString: Constructor for UnicodeString Object
Arguments:
inputString - The string to convert to UnicodeString (does not mutate)
Throws:
InvalidUTF8Exception - if the inputString is not proper UTF-8
*/
UnicodeString::UnicodeString( std::string inputString )
{
  std::string::iterator it = inputString.begin();
  // Iterate over inputString and store each unicode Character as a string in a vector
  while ( it != inputString.end() )
  {
    // Test if Character is 1,2,3 or 4 bytes; throw exception if invalid
    if ( ( ( *it | 0x7F) & 0x80 ) == 0x00 )
    {
      std::string charString;
      charString.append( 1, *it );
      uniString.push_back( charString );
      it++;
    }
    else if ( ( ( *it | 0x1F) & 0xE0 ) == 0xC0 )
    {
      std::string charString;
      charString.append( 1, *it );
      it++;
      // Check Following Character
      if ( ( ( *it | 0x3F ) & 0xC0 ) == 0x80 )
      {
        charString.append( 1, *it );
        uniString.push_back( charString );
        it++;
      }
      else
      {
        throw InvalidUTF8Exception( inputString );
      }
    }
    else if ( ( ( *it | 0x0F) & 0xF0 ) == 0xE0 )
    {
      std::string charString;
      charString.append( 1, *it );
      it++;
      // Iterate over next two characters
      for ( int i = 0; i < 2; i++ )
      {
        // Check if character is valid
        if ( ( ( *it | 0x3F ) & 0xC0 ) == 0x80 )
        {
          charString.append( 1, *it );
          it++;
        }
        else
        {
          throw InvalidUTF8Exception( inputString );
        }
      }
      uniString.push_back( charString );
    }
    else if ( ( (*it | 0x07) & 0xF8 ) == 0xF0 )
    {
      std::string charString;
      charString.append( 1, *it );
      it++;
      // Iterate over next 3 characters
      for ( int i = 0; i < 3; i++ )
      {
        // Check if character is valid
        if ( ( ( *it | 0x3F ) & 0xC0 ) == 0x80 )
        {
          charString.append( 1, *it );
          it++;
        }
        else
        {
          throw InvalidUTF8Exception( inputString );
        }
      }
      uniString.push_back( charString );
    }
    else
    {
      throw InvalidUTF8Exception( inputString );
    }
  }
}

/*
toString: Returns a UnicodeString object as a regular string
Arguments: None
Returns: String version of UnicodeString
*/
std::string UnicodeString::toString()
{
  std::string outputString;
  // Iterate through string and build up output string
  for ( std::vector<std::string>::iterator it = uniString.begin(); it != uniString.end(); it++ )
  {
    outputString.append(*it);
  }
  return outputString;
}

void UnicodeString::insert( int insertLocation, UnicodeString stringToInsert )
{
  std::vector<std::string> insertVector = stringToInsert.getVector();
  uniString.insert( unString.begin() + insertLocation, insertVector.begin(), insertVector.end() );
  return;
}

void UnicodeString::append( UnicodeString stringToAppend )
{
  uniString.insert( uniString.end(), stringToAppend.begin(), stringToAppend.end() );
  return;
}
