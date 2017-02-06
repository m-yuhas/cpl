// Invalid UTF-8 String Exception
// (C) 2016 Michael Yuhas

#pragma once
#include <string>

/*
InvalidUTF8Exception: Handles Exceptions Caused When a String Is not Proper UTF-8
Public Methods:
Constructor - constructs an InvalidUTF8Exception Object
what() - returns a string describing what went wrong
Protected Attributes:
badString - the offending malformed UTF-8 string
*/
class InvalidUTF8Exception
{
  public:
    InvalidUTF8Exception( std::string badString );
    std::string what();
  protected:
    std::string badString;
};
