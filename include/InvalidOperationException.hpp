// Invalid Operation Exception
// (C) 2016 Michael Yuhas

#pragma once
#include <string>

/*
InvalidOperationException: Handles Exceptions Caused When Attempting an Operation on incompatible data types
Public Methods:
Constructor - constructs an InvalidOperationException Object
what() - returns a string describing what went wrong
Protected Attributes:
op - the offending Operation
arg1 - the type of the first argument
arg2 - the type of the second argument
*/
class InvalidOperationException
{
  public:
    InvalidOperationException( char op, char arg1, char arg2 );
    std::string what();
  protected:
    char op;
    char arg1;
    char arg2;
};
