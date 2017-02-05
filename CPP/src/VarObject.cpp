// This class defines an Object to Store a Variable
// (C) 2016 Michael Yuhas

#include <iostream>
#include "../include/VarObject.hpp"
//#include "../include/UnicodeString.hpp"
//#include "../include/StringUtils.hpp"
//s#include "../include/InvalidOperationException.hpp"


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

VarObject::VarObject( std::vector<VarObject> array )
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

void VarObject::setValue( UnicodeString ustr )
{
  uVal = ustr;
  type = 4;
}

void VarObject::setValue( std::vector<VarObject> array )
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

std::vector<VarObject> VarObject::getArrayVal()
{
  return arrVal;
}

VarObject VarObject::add( VarObject addend )
{
  switch( type )
  {
    case 1 :
    {
      switch( addend.getType() )
      {
        case 1 :
        {
          return VarObject( bVal | addend.getBoolVal() );
        }
        case 2 :
        {
          return VarObject( (int)bVal + addend.getIntVal() );
        }
        case 3 :
        {
          return VarObject( (double)bVal + addend.getDoubleVal() );
        }
        case 4 :
        {
          if ( bVal )
          {
            UnicodeString returnString = addend.getUStringVal();
            returnString.insert( 0, UnicodeString("是") );
            return VarObject( returnString );
          }
          else
          {
            UnicodeString returnString = addend.getUStringVal();
            returnString.insert( 0, UnicodeString("否") );
            return VarObject( returnString );
          }
        }
        case 5 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = addend.getArrayVal().begin(); it != addend.getArrayVal().end(); it++ )
          {
            try
            {
              newArray.push_back( this->add( *it ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        }
        default :
        {
          throw InvalidOperationException( 1, type, addend.getType() );
        }
      }
    }
    case 2 :
    {
      switch( addend.getType() )
      {
        case 1 :
        {
          return VarObject( iVal + (int)addend.getBoolVal() );
        }
        case 2 :
        {
          return VarObject( iVal + addend.getIntVal() );
        }
        case 3 :
        {
          return VarObject( (double)iVal + addend.getDoubleVal() );
        }
        case 4 :
        {
          return VarObject( std::to_string( iVal ) + addend.getUStringVal().toString() );
        }
        case 5 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = addend.getArrayVal().begin(); it != addend.getArrayVal().end(); it++ )
          {
            try
            {
              newArray.push_back( this->add( *it ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        }
        default:
        {
          throw InvalidOperationException( 1, type, addend.getType() );
        }
      }
    }
    case 3 :
    {
      switch( addend.getType() )
      {
        case 1 :
        {
          return VarObject( dVal + (double)addend.getBoolVal() );
        }
        case 2 :
        {
          return VarObject( dVal + (double)addend.getIntVal() );
        }
        case 3 :
        {
          return VarObject( dVal + addend.getDoubleVal() );
        }
        case 4 :
        {
          return VarObject( std::to_string( dVal ) + addend.getUStringVal().toString() );
        }
        case 5 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = addend.getArrayVal().begin(); it != addend.getArrayVal().end(); it++ )
          {
            try
            {
              newArray.push_back( this->add( *it ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
        }
        default:
        {
          throw InvalidOperationException( 1, type, addend.getType() );
        }
      }
    }
    case 4 :
    {
      switch( addend.getType() )
      {
        case 1 :
        {
          if ( addend.getBoolVal() )
          {
            UnicodeString returnString = addend.getUStringVal();
            returnString.insert( 0, UnicodeString("是") );
            return VarObject( returnString );
          }
          else
          {
            UnicodeString returnString = addend.getUStringVal();
            returnString.insert( 0, UnicodeString("否") );
            return VarObject( returnString );
          }
        }
        case 2 :
        {
          UnicodeString returnString = uVal;
          returnString.append( UnicodeString( std::to_string( addend.getIntVal() ) ) );
          return VarObject( returnString );
        }
        case 3 :
        {
          UnicodeString returnString = uVal;
          returnString.append( UnicodeString( std::to_string( addend.getDoubleVal() ) ) );
          return VarObject( returnString );
        }
        case 4 :
        {
          UnicodeString returnString = uVal;
          returnString.append( addend.getUStringVal() );
          return VarObject( returnString );
        }
        case 5 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = addend.getArrayVal().begin(); it != addend.getArrayVal().end(); it++ )
          {
            try
            {
              newArray.push_back( this->add( *it ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        }
        default:
        {
          throw InvalidOperationException( 1, type, addend.getType() );
        }
      }
    }
    case 5 :
    {
      switch( addend.getType() )
      {
        case 1 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
          {
            try
            {
              newArray.push_back( (*it).add( addend.getBoolVal() ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        }
        case 2 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
          {
            try
            {
              newArray.push_back( (*it).add( addend.getIntVal() ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        }
        case 3 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
          {
            try
            {
              newArray.push_back( (*it).add( addend.getDoubleVal() ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        }
        case 4 :
        {
          std::vector<VarObject> newArray;
          for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
          {
            try
            {
              newArray.push_back( (*it).add( addend.getUStringVal() ) );
            }
            catch ( InvalidOperationException e )
            {
              throw e;
            }
          }
          return VarObject( newArray );
        }
        case 5 :
        {
          std::vector<VarObject> returnArray = arrVal;
          returnArray.insert(arrVal.end(), addend.getArrayVal().begin(), addend.getArrayVal().end());
          return VarObject( returnArray );
        }
        default :
        {
          throw InvalidOperationException( 1, type, addend.getType() );
        }
      }
    }
    default :
    {
      throw InvalidOperationException( 1, type, addend.getType() );
    }
  }
  throw InvalidOperationException( 1, type, addend.getType() );
}

VarObject VarObject::sub( VarObject var )
{
  //TODO: Implement This
  return VarObject( false );
}


bool VarObject::equals( VarObject var )
{
  switch( type )
  {
    case 1 :
    {
      switch( var.getType() )
      {
        case 1 :
        {
          return bVal == var.getBoolVal();
        }
        case 2 :
        {
          return (int)bVal == var.getIntVal();
        }
        case 3 :
        {
          return (double)bVal == var.getDoubleVal();
        }
        case 4 :
        {
          if ( bVal )
          {
            if ( std::string("是").compare( var.getUStringVal().toString() ) == 0 )
            {
              return true;
            }
            else
            {
              return false;
            }
          }
          else
          {
            if ( std::string("否").compare( var.getUStringVal().toString() ) == 0 )
            {
              return true;
            }
            else
            {
              return false;
            }
          }
        }
        case 5 :
        {
          if ( var.getArrayVal().size() == 1 )
          {
            return this->equals(var.getArrayVal().front());
          }
          else
          {
            return false;
          }
        }
        default:
        {
          throw InvalidOperationException( 101, type, var.getType() );
        }
      }
    }
    case 2 :
    {
      switch( var.getType() )
      {
        case 1 :
        {
          return iVal == (int)var.getBoolVal();
        }
        case 2 :
        {
          return iVal == var.getIntVal();
        }
        case 3 :
        {
          return (double)iVal == var.getDoubleVal();
        }
        case 4 :
        {
          if ( std::to_string( iVal ).compare( var.getUStringVal().toString() ) == 0 )
          {
            return true;
          }
          else
          {
            return false;
          }
        }
        case 5 :
        {
          if ( var.getArrayVal().size() == 1 )
          {
            return this->equals(var.getArrayVal().front());
          }
          else
          {
            return false;
          }
        }
        default:
        {
          throw InvalidOperationException( 101, type, var.getType() );
        }
      }
    }
    case 3 :
    {
      switch( var.getType() )
      {
        case 1 :
        {
          return dVal == (double)var.getBoolVal();
        }
        case 2 :
        {
          return dVal == (double)var.getIntVal();
        }
        case 3 :
        {
          return dVal == var.getDoubleVal();
        }
        case 4 :
        {
          if ( std::to_string( dVal ).compare( var.getUStringVal().toString() ) == 0 )
          {
            return true;
          }
          else
          {
            return false;
          }
        }
        case 5 :
        {
          if ( var.getArrayVal().size() == 1 )
          {
            return this->equals(var.getArrayVal().front());
          }
          else
          {
            return false;
          }
        }
        default:
        {
          throw InvalidOperationException( 101, type, var.getType() );
        }
      }
    }
    case 4 :
    {
      switch( var.getType() )
      {
        case 1 :
        {
          if ( var.getBoolVal() )
          {
            if ( uVal.toString().compare( std::string("是") ) == 0 )
            {
              return true;
            }
            else
            {
              return false;
            }
          }
          else
          {
            if ( uVal.toString().compare( std::string("否") ) == 0 )
            {
              return true;
            }
            else
            {
              return false;
            }
          }
        }
        case 2 :
        {
          if ( uVal.toString().compare( std::to_string( var.getIntVal() ) ) == 0 )
          {
            return true;
          }
          else
          {
            return false;
          }
        }
        case 3 :
        {
          if ( uVal.toString().compare( std::to_string( var.getDoubleVal() ) ) == 0 )
          {
            return true;
          }
          else
          {
            return false;
          }
        }
        case 4 :
        {
          if ( uVal.toString().compare( var.getUStringVal().toString() ) == 0 )
          {
            return true;
          }
          else
          {
            return false;
          }
        }
        case 5 :
        {
          if ( var.getArrayVal().size() == 1 )
          {
            return this->equals(var.getArrayVal().front());
          }
          else
          {
            return false;
          }
        }
        default:
        {
          throw InvalidOperationException( 101, type, var.getType() );
        }
      }
    }
    case 5 :
    {
      switch( var.getType() )
      {
        case 1 :
        {
          if( arrVal.size() == 1 )
          {
            return arrVal.front().equals(var);
          }
          else
          {
            return false;
          }
        }
        case 2 :
        {
          if( arrVal.size() == 1 )
          {
            return arrVal.front().equals(var);
          }
          else
          {
            return false;
          }
        }
        case 3 :
        {
          if( arrVal.size() == 1 )
          {
            return arrVal.front().equals(var);
          }
          else
          {
            return false;
          }
        }
        case 4 :
        {
          if( arrVal.size() == 1 )
          {
            return arrVal.front().equals(var);
          }
          else
          {
            return false;
          }
        }
        case 5 :
        {
          std::cout << "HERE";
          if ( arrVal.size() == var.getArrayVal().size() )
          {
            std::vector<VarObject>::iterator it2 = var.getArrayVal().begin();
            for ( std::vector<VarObject>::iterator it = arrVal.begin(); it != arrVal.end(); it++ )
            {
              std::cout << "Array Member Type 1 " << std::to_string((*it).getType());
              std::cout << "Array Member Type 2 " << std::to_string((*it2).getType());
              if ( ! (*it).equals( *it2 ) )
              {
                return false;
              }
              it2++;
            }
            return true;
          }
          else
          {
            return false;
          }
        }
        default :
        {
          throw InvalidOperationException( 101, type, var.getType() );
        }
      }
    }
    default :
    {
      throw InvalidOperationException( 101, type, var.getType() );
    }
  }
  throw InvalidOperationException( 101, type, var.getType() );
}
