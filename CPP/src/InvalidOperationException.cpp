// Invalid Operation Exception
// (C) 2016 Michael Yuhas

#include <string>
#include "../include/InvalidOperationException.hpp"

/*
Exeception Constructor: Constructs an InvalidOperationException Object
Arguments:
optype - 1=add, 2=sub, 3=mul, 4=div, 5=mod, 6=exp, 7=fac
optype - 101=equals, 102=not equals, 103=greater than, 104=less than, 105=gte, 106=lte
arg1type - data type of caller object: 1=bool, 2=int, 3=double, 4=string, 5=array
arg2type - data type of callee object: 1=bool, 2=int, 3=double, 4=string, 5=array
*/
InvalidOperationException::InvalidOperationException( char optype, char arg1type, char arg2type )
{
  op = optype;
  arg1 = arg1type;
  arg2 = arg2type;
}

/*
what(): Returns a description of what went wrong
Arguments: None
Returns: Error Message
*/
std::string InvalidUTF8Exception::what()
{
  std::string opString;
  std::string arg1String;
  std::string arg2String;
  // Parse Operation Error Code
  switch( op )
  {
    case 1 :
      opString = "+";
      break;
    case 2 :
      opString = "-";
      break;
    case 3 :
      opString = "*";
      break;
    case 4 :
      opString = "/";
      break;
    case 5 :
      opString = "%";
      break;
    case 6 :
      opString = "^";
      break;
    case 7 :
      opString = "!";
      break;
    case 101 :
      opString = "=";
      break;
    case 103 :
      opString = ">";
      break;
    case 104 :
      opString = "<";
      break;
    default :
      opString = "Unknown Operation Type";
      break;
  }
  // Parse arg1 error code
  switch( arg1 )
  {
    case 1 :
      arg1String = "Boolean";
      break;
    case 2 :
      arg1String = "Integer";
      break;
    case 3 :
      arg1String = "Double";
      break;
    case 4 :
      arg1String = "String";
      break;
    case 5 :
      arg1String = "Array";
      break;
    default :
      arg1String = "Unknown Type";
      break;
  }
  // Parse arg2 error code
  switch( arg2 )
  {
    case 1 :
      arg2String = "Boolean";
      break;
    case 2 :
      arg2String = "Integer";
      break;
    case 3 :
      arg2String = "Double";
      break;
    case 4 :
      arg2String = "String";
      break;
    case 5 :
      arg2String = "Array";
      break;
    default :
      arg2String = "Unknown Type";
      break;
  }
  return "Invalid Operation: \"" + opString + "\" on \"" + arg1String + "\" and \"" + arg2String + "\"";
}
