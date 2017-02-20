// This class includes various parsers for mathematical expressions
// (C) 2016 Michael Yuhas

#pragma once
#include "../include/Parser.hpp"

VarObject Parser::evaluateExpression( UnicodeString expression )
{
  unsigned char parenthCount = 0;
  unsigned char addSubIndex = -1;
  unsigned char mulDivIndex = -1;
  unsigned char expIndex = -1;
  unsigned char facIndex = -1;
  unsigned char optype = -1;
  for ( int i = 0; i <= expression.length(); i++ )
  {
    if ( expression.getChar(i).compare("(") == 0 )
    {
      parenthCount += 1;
    }
    else if ( expression.getChar(i).compare("(") == 0 )
    {
      parenthCount -= 1;
    }
    else if ( parenthCount == 0 )
    {
      if ( expression.getChar(i).compare("+") == 0 )
      {
        addSubIndex = currIndex;
        optype = 1;
        break;
      }
      else if ( expression.getChar(i).compare("-") == 0 )
      {
        addSubIndex = currIndex;
        optype = 2;
        break;
      }
      else if ( expression.getChar(i).compare("*") == 0 )
      {
        mulDivIndex = currIndex;
        optype = 3;
        break;
      }
      else if ( expression.getChar(i).compare("/") == 0 )
      {
        mulDivIndex = currIndex;
        optype = 4;
        break;
      }
      else if ( expression.getChar(i).compare("%") == 0 )
      {
        mulDivIndex = currIndex;
        optype = 5;
        break;
      }
      else if ( expression.getChar(i).compare("^") == 0 )
      {
        expIndex = currIndex;
        optype = 6;
        break;
      }
      else if ( expression.getChar(i).compare("!") == 0 )
      {
        facIndex = currIndex;
        optype = 7;
        break;
      }
    }
    currIndex++;
  }

  if ( addSubIndex != -1 )
  {
    UnicodeString upper = expression.substring( 0, addSubIndex );
    UnicodeString lower = expression.substring( addSubIndex+1, expression.length() );
    if ( optype == 1 )
    {
      return Parser::evaluateExpression( upper ).add( Parser::evaluateExpression( lower ) );
    }
    if ( optype == 2 )
    {
      return Parser::evaluateExpression( upper ).sub( Parser::evaluateExpression( lower ) );
    }
  }
  if ( mulDivIndex != -1 )
  {
    UnicodeString upper = expression.substring( 0, mulDivIndex );
    UnicodeString lower = expression.substring( mulDivIndex+1, expression.length() );
    if ( optype == 3 )
    {
      return Parser::evaluateExpression( upper ).mul( Parser::evaluateExpression( lower ) );
    }
    if ( optype == 4 )
    {
      return Parser::evaluateExpression( upper ).div( Parser::evaluateExpression( lower ) );
    }
    if ( optype == 5 )
    {
      return Parser::evaluateExpression( upper ).mod( Parser::evaluateExpression( lower ) );
    }
  }
  if ( expIndex != -1 && optype == 6 )
  {
    UnicodeString upper = expression.substring( 0, expIndex );
    return Parser::evaluateExpression( upper ).exp( Parser::evaluateExpression( lower ) );
  }
  if ( facIndex != -1 && optype == 7 )
  {
    return Parser::evaluateExpression( upper ).fac();
  }
  return Parser::evaluateAtom( expression );
}

VarObject Parser::evaluateExpressionBoolean( UnicodeString expression )
{

}

VarObject Parser::evaluateAtom( UnicodeString expression )
{
  if ( expression.length() <= 0 )
  {
    throw
  }
  if ( expression.toString().compare("是") == 0 )
  {
    return VarObject( true );
  }
  if ( expression.toString().compare("否") == 0 )
  {
    return VarObject( false );
  }
  try
  {
    return VarObject( std::stoi( expression.toString() ) );
  }
  catch ( invalid_argument )
  {
    try
    {
      //TODO: Search Hashmap for Instance of Var.
    }
    catch
    {

    }
  }
}
