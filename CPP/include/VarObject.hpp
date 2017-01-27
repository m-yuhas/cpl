// This class defines an Object to Store a Variable
// (C) 2016 Michael Yuhas

#include <vector>
#include "../include/UnicodeString.hpp"

class VarObject
{
  public:
    VarObject( bool b );
    VarObject( int i );
    VarObject( double d );
    VarObject( UnicodeString ustr );
    VarObject( std::vector<std::VarObject> array );
    void setValue( bool b );
    void setValue( int i );
    void setValue( double d );
    void setValue( UnicodeString ustr );
    void setValue( std::vector<std::VarObject> array );
    char getType();
    bool getBoolVal();
    int getIntVal();
    double getDoubleVal();
    UnicodeString getUStringVal();
    std::vector<std::VarObject> getArrayVal();
    VarObject add( VarObject addend );
    VarObject sub( VarObject subtrahend );
    VarObject mul( VarObject factor );
    VarObject div( VarObject divisor );
    VarObject mod( VarObject divisor );
    VarObject exp( VarObject exponent );
    VarObject fac();
    bool equals( VarObject var );
    bool notEquals( VarObject var );
    bool greaterThan( VarObject var );
    bool lessThan( VarObject var );
    bool greaterThanOrEquals( VarObject var );
    bool lessThanOrEquals( VarObject var );
  protected:
    char type;
    bool bVal;
    int iVal;
    double dVal;
    UnicodeString uVal;
    std::vector<std::VarObject> arrVal;
}
