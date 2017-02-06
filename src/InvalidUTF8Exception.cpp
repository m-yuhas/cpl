// Invalid UTF-8 String Exception
// (C) 2016 Michael Yuhas

#include <string>
#include "../include/InvalidUTF8Exception.hpp"

/*
Exeception Constructor: Constructs an InvalidUTF8Exception Object
Arguments:
str - the offending malformed UTF-8 String
*/
InvalidUTF8Exception::InvalidUTF8Exception( std::string str )
{
  badString = str;
}

/*
what(): Returns a description of what went wrong
Arguments: None
Returns: Error Message
*/
std::string InvalidUTF8Exception::what()
{
  return "Invalid UTF-8 String: \"" + badString + "\"";
}
