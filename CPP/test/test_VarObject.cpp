#include <gtest/gtest.h>
#include "../include/VarObject.hpp"

TEST( VarObject, BooleanConstructor )
{
  VarObject boolObj( true );
  VarObject boolObj2( true );
  EXPECT_EQ( true, boolObj.equals( boolObj2 ) );
}

TEST( VarObject, IntegerConstructor )
{
  VarObject intObj( 2 );
  VarObject intObj2( 2 );
  EXPECT_EQ( true, intObj.equals( intObj2 ) );
}

TEST( VarObject, DoubleConstructor )
{
  VarObject doubleObj( 2.1 );
  VarObject doubleObj2( 2.1 );
  EXPECT_EQ( true, doubleObj.equals( doubleObj2 ) );
}

TEST( VarObject, StringConstructor )
{
  VarObject stringObj( UnicodeString("你好世界！") );
  VarObject stringObj2( UnicodeString("你好世界！") );
  EXPECT_EQ( true, stringObj.equals( stringObj2 ) );
}
/*
TEST( VarObject, ArrayConstructor )
{
  std::vector<VarObject> arr;
  arr.push_back( VarObject( true ) );
  VarObject arrObj( arr );
  VarObject arrObj2( arr );
  EXPECT_EQ( true, arrObj.equals( arrObj2 ) );
}*/
TEST( VarObject, AddBooleanToBoolean )
{
  VarObject Obj1( true );
  VarObject Obj2( false );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( true ) ) );
}

TEST( VarObject, AddBooleanToInt )
{
  VarObject Obj1( true );
  VarObject Obj2( 1 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 2) ) );
}

TEST( VarObject, AddBooleanToDouble )
{
  VarObject Obj1( true );
  VarObject Obj2( 4.5 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 5.5)));
}

TEST( VarObject, AddBooleanToString )
{
  VarObject Obj1( true );
  VarObject Obj2( UnicodeString("Hello") );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( UnicodeString("是Hello"))));
}

TEST( VarObject, AddIntToBoolean )
{
  VarObject Obj1( 10 );
  VarObject Obj2( false );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 10 )));
}

TEST( VarObject, AddIntToInt )
{
  VarObject Obj1( 11 );
  VarObject Obj2( 12 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 23 )));
}

TEST( VarObject, AddIntToDouble )
{
  VarObject Obj1( 13 );
  VarObject Obj2( 13.12 );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 26.12 )));
}

TEST( VarObject, AddIntToString )
{
  VarObject Obj1( 14 );
  VarObject Obj2( UnicodeString("Hello") );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( UnicodeString("14Hello"))));
}

TEST( VarObject, AddDoubleToBoolean )
{
  VarObject Obj1( 15.51 );
  VarObject Obj2( true );
  EXPECT_EQ( true, Obj1.add(Obj2).equals( VarObject( 16.61 )));
}
