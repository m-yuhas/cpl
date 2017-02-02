#include <gtest/gtest.h>
#include "../include/VarObject.hpp"

TEST( VarObject, BooleanConstructor )
{
  VarObject boolObj( true );
  VarObject boolObj2( true );
  EXPECT_EQ( true, boolObj.equals( boolObj2 ) );
}
