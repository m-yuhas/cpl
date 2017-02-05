#include <gtest/gtest.h>
#include "../include/VarObject.hpp"

TEST( VarObjectConstructors, BooleanConstructor )
{
  VarObject boolObj( true );
  VarObject boolObj2( true );
  EXPECT_EQ( true, boolObj.equals( boolObj2 ) );
}

TEST( VarObjectConstructors, IntegerConstructor )
{
  VarObject intObj( 2 );
  VarObject intObj2( 2 );
  EXPECT_EQ( true, intObj.equals( intObj2 ) );
}

TEST( VarObjectConstructors, DoubleConstructor )
{
  VarObject doubleObj( 2.1 );
  VarObject doubleObj2( 2.1 );
  EXPECT_EQ( true, doubleObj.equals( doubleObj2 ) );
}

TEST( VarObjectConstructors, StringConstructor )
{
  VarObject stringObj( UnicodeString("你好世界！") );
  VarObject stringObj2( UnicodeString("你好世界！") );
  EXPECT_EQ( true, stringObj.equals( stringObj2 ) );
}
/*
TEST( VarObjectConstructors, ArrayConstructor )
{
  std::vector<VarObject> arr;
  arr.push_back( VarObject( true ) );
  VarObject arrObj( arr );
  VarObject arrObj2( arr );
  try
  {
    EXPECT_EQ( true, arrObj.equals( arrObj2 ) );
  }
  catch( InvalidOperationException )
  {
    FAIL() << "Wrong Exception";
  }
}
*/
TEST( VarObjectAddMethod, AddBooleanToBoolean )
{
  VarObject Obj1( true );
  VarObject Obj2( false );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( true ) ) );
}

TEST( VarObjectAddMethod, AddBooleanToInt )
{
  VarObject Obj1( true );
  VarObject Obj2( 1 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 2) ) );
}

TEST( VarObjectAddMethod, AddBooleanToDouble )
{
  VarObject Obj1( true );
  VarObject Obj2( 4.5 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 5.5)));
}

TEST( VarObjectAddMethod, AddBooleanToString )
{
  VarObject Obj1( true );
  VarObject Obj2( UnicodeString("Hello") );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( UnicodeString("是Hello"))));
}

TEST( VarObjectAddMethod, AddIntToBoolean )
{
  VarObject Obj1( 10 );
  VarObject Obj2( false );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 10 )));
}

TEST( VarObjectAddMethod, AddIntToInt )
{
  VarObject Obj1( 11 );
  VarObject Obj2( 12 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 23 )));
}

TEST( VarObjectAddMethod, AddIntToDouble )
{
  VarObject Obj1( 13 );
  VarObject Obj2( 13.12 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 26.12 )));
}

TEST( VarObjectAddMethod, AddIntToString )
{
  VarObject Obj1( 14 );
  VarObject Obj2( UnicodeString("Hello") );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( UnicodeString("14Hello"))));
}

TEST( VarObjectAddMethod, AddDoubleToBoolean )
{
  VarObject Obj1( 15.51 );
  VarObject Obj2( true );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 16.61 )));
}

TEST( VarObjectSubtractMethod, SubtractBooleanFromBoolean )
{
  VarObject Obj1( true );
  VarObject Obj2( false );
  try
  {
    Obj1.sub(Obj2);
    FAIL() << "Expected Invalid Operation Exception";
  }
  catch( InvalidOperationException err)
  {
    EXPECT_EQ( err.what(), "Invalid Operation: \"-\" on \"Boolean\" and \"Boolean\"");
  }
  catch(...)
  {
    FAIL() << "Expected Invalid Operation Exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractIntFromBoolean )
{
  VarObject Obj1( true );
  VarObject Obj2( 1 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( 0 )));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractDoubleFromBoolean )
{
  VarObject Obj1( true );
  VarObject Obj2( 1.1 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( -0.1)));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractStringFromBooleanStringContained )
{
  VarObject Obj1( false );
  VarObject Obj2( UnicodeString("否"));
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString( "" ))));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractStringFromBooleanStringNotContained )
{
  VarObject Obj1( false );
  VarObject Obj2( UnicodeString("字符串"));
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString( "否" ))));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractBooleanFromInt )
{
  VarObject Obj1( 12 );
  VarObject Obj2( true );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( 11)));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractIntFromInt )
{
  VarObject Obj1( 24 );
  VarObject Obj2( 48 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( -24)));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractDoubleFromInt )
{
  VarObject Obj1( 48 );
  VarObject Obj2( 35.4 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( 12.6)));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractStringFromIntStringIncluded )
{
  VarObject Obj1( 101 );
  VarObject Obj2( UnicodeString( "0"));
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString("11"))));
  }
  catch(...)
  {
    FAIL() << "Did not expect excpetion";
  }
}

TEST( VarObjectSubtractMethod, SubtractStringFromIntStringNotIncluded )
{
  VarObject Obj1( 101 );
  VarObject Obj2( UnicodeString( "2"));
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString("101"))));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractBooleanFromDouble )
{
  VarObject Obj1( 1.2345 );
  VarObject Obj2( true );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( 0.2345 )));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractIntFromDouble )
{
  VarObject Obj1( 1.2345 );
  VarObject Obj2( 12 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( -10.8765)));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractDoubleFromDouble )
{
  VarObject Obj1( 1.2345 );
  VarObject Obj2( 1.2345 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( 0.0 )));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractStringFromDouble )
{
  VarObject Obj1( 1.2345 );
  VarObject Obj2( UnicodeString("."));
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( "12345")));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractBooleanFromString )
{
  VarObject Obj1( UnicodeString("字符串是不是包括"));
  VarObject Obj2( true );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString("字符串不包括"))));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractIntFromString )
{
  VarObject Obj1( UnicodeString("This is a string with numbers 123456"));
  VarObject Obj2( 1 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString("This is a string with numbers 23456"))));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractDoubleFromString )
{
  VarObject Obj1( UnicodeString("This is a string with numbers 123456"));
  VarObject Obj2( 1.1 );
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString("This is a string with numbers 123456"))));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}

TEST( VarObjectSubtractMethod, SubtractStringFromString )
{
  VarObject Obj1( UnicodeString("字符串"));
  VarObject Obj2( UnicodeString("字"));
  try
  {
    EXPECT_EQ( true, Obj1.sub(Obj2).equals( VarObject( UnicodeString("符串"))));
  }
  catch(...)
  {
    FAIL() << "Did not expect exception";
  }
}
