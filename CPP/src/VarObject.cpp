// This class defines an Object to Store a Variable
// (C) 2016 Michael Yuhas

#include "../include/VarObject.hpp"
#include "../include/UnicodeString.hpp"

VarObject::VarObject( bool b )
{
  bVal = b;
  type = 1;
}

VarObject::VarObject( int i )
{
  iVal = i;
  type = 2;
}

VarObject::VarObject( double d )
{
  dVal = d;
  type = 3;
}

VarObject::VarObject( UnicodeString ustr )
{
  uVal = ustr;
  type = 4;
}

VarObject::VarObject( std::vector<std::VarObject> array )
{
  arrVal = array;
  type = 5;
}

void VarObject::setValue( bool b )
{
  bVal = b;
  type = 1;
}

void VarObject::setValue( int i )
{
  iVal = i;
  type = 2;
}

void VarObject::setValue( double d )
{
  dVal = d;
  type = 3;
}

void VarObject::setValue( UnicodeString ustr )
{
  uVal = ustr;
  type = 4;
}

void VarObject::setValue( std::vector<std::VarObject> array )
{
  arrVal = array;
  type = 5;
}

char VarObject::getType()
{
  return type;
}

bool VarObject::getBoolVal()
{
  return bVal;
}

int VarObject::getIntVal()
{
  return iVal;
}

double VarObject::getDoubleVal()
{
  return dVal;
}

UnicodeString VarObject::getUStringVal()
{
  return uVal;
}

std::vector<std::VarObject> VarObject::getArrayVal()
{
  return arrVal;
}

VarObject add( VarObject addend )
{
  switch( type )
  {
    case 1 :
      switch( addend.getType() )
      {
        case 1 :
          return VarObject( bVal | addend.getBoolVal() );
        case 2 :
          return VarObject( (int)bVal + addend.getIntVal() );
        case 3 :
          return VarObject( (double)bVal + addend.getDoubleVal() );
        case 4 :
          if ( bVal )
          {
            return VarObject( addend.uVal.insert( 0, UnicodeString("是") ) );
          }
          else
          {
            return VarObject( addend.uVal.insert( 0, UnicodeString("否") ) );
          }
        case 5 :
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
          {
            try
            {
              newArray.push_back( this.add( *it ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        default:
          throw InvalidOperationException( 1, type, addend.getType() )
      }
    case 2 :
      switch( addend.getType() )
      {
        case 1 :
          return VarObject( iVal + (int)addend.getBoolVal() );
        case 2 :
          return VarObject( iVal + addend.getIntVal() );
        case 3 :
          return VarObject( (double)iVal + addend.getDoubleVal() );
        case 4 :
          return VarObject( std::to_string( iVal ) + addend.getStringValue() );
        case 5 :
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
          {
            try
            {
              newArray.push_back( this.add( *it ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        default:
          throw InvalidOperationException( 1, type, addend.getType() )
      }
    case 3 :
      switch( addend.getType() )
      {
        case 1 :
          return VarObject( dVal + (double)addend.getBoolVal() );
        case 2 :
          return VarObject( dVal + (double)addend.getIntVal() );
        case 3 :
          return VarObject( dVal + addend.getDoubleVal() );
        case 4 :
          return VarObject( std::to_string( dVal ) + addend.getStringValue() );
        case 5 :
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
          {
            try
            {
              newArray.push_back( this.add( *it ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        default:
          throw InvalidOperationException( 1, type, addend.getType() )
      }
    case 4 :
      break;
    case 5 :
      break;
    default:
      throw InvalidOperationException( 1, type, addend.getType() )
  }
  return;
}
