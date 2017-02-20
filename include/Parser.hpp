// This class includes various parsers for mathematical expressions
// (C) 2016 Michael Yuhas

#pragma once
#include "../include/InvalidOperationException.hpp"
#include "../include/VarObject.hpp"
#include "../include/UnicodeString.hpp"

class Parser
{
  public:
    static VarObject evaluateExpression( UnicodeString expression );
    static VarObject Parser::evaluateExpressionBoolean( UnicodeString expression );
  private:
    static VarObject evaluateAtom( UnicodeString expression );
};
