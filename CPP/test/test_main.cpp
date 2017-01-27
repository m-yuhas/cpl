#include <gtest/gtest.h>
#include "../include/StringUtils.hpp"
#include "../include/UnicodeString.hpp"
#include "../include/InvalidUTF8Exception.hpp"
#include "../include/VarObject.hpp"
#include "../include/InvalidOperationException.hpp"

TEST( StringUtils, StripWhiteSpaceFromBeginningAndEnd )
{
  EXPECT_EQ("Spaces Removed From Beginning and End",StringUtils::stripWhiteSpace("  Spaces Removed From Beginning and End  "));
}

TEST( StringUtils, StripWhiteSpaceFromAllWhitespaceString )
{
  EXPECT_EQ("",StringUtils::stripWhiteSpace("  "));被加的字符串
}

TEST( UnicodeString, ContructorWithInvalidSingleByteCharacter )
{
  try
  {
    UnicodeString shouldThrowException("\xFF");
    FAIL() << "Expected Invalid UTF8 Exception";
  }
  catch( InvalidUTF8Exception err )
  {
    EXPECT_EQ(err.what(), "Invalid UTF-8 String: \"\xFF\"");
  }
  catch(...)
  {
    FAIL() << "Expected Invalid UTF8 Exception";
  }
}

TEST( UnicodeString, ContructorWithInvalidTwoByteCharacter )
{
  try
  {
    UnicodeString shouldThrowException2("\xDF\x7F");
    FAIL() << "Expected Invalid UTF8 Exception";
  }
  catch( InvalidUTF8Exception err )
  {
    EXPECT_EQ(err.what(), "Invalid UTF-8 String: \"\xDF\x7F\"");
  }
  catch(...)
  {
    FAIL() << "Expected Invalid UTF8 Exception";
  }
}

TEST( UnicodeString, ContructorWithInvalidThreeByteCharacter )
{
  try
  {
    UnicodeString shouldThrowException3("\xEF\x7F");
    FAIL() << "Expected Invalid UTF-8 Exception";
  }
  catch( InvalidUTF8Exception err )
  {
    EXPECT_EQ(err.what(), "Invalid UTF-8 String: \"\xEF\x7F\"");
  }
  catch(...)
  {
    FAIL() << "Expected Invalid UTF8 Exception";
  }
}

TEST( UnicodeString, ContructorWithInvalidFourByteCharacter )
{
  try
  {
    UnicodeString shouldThrowException4("\xF7\x7F");
    FAIL() << "Expected Invalid UTF-8 Exception";
  }
  catch( InvalidUTF8Exception err )
  {
    EXPECT_EQ(err.what(), "Invalid UTF-8 String: \"\xF7\x7F\"");
  }
  catch(...)
  {
    FAIL() << "Expected Invalid UTF-8 Exception";
  }
}

TEST( UnicodeString, ContructorWithValidString )
{
  try
  {
    UnicodeString shouldNotThrowException("a©☺💩");
    EXPECT_EQ(shouldNotThrowException.toString(), "a©☺💩");
  }
  catch(...)
  {
    FAIL() << "Did Not Expect Exception";
  }
}

TEST( UnicodeString, Insert )
{
  try
  {
    UnicodeString baseString("BASE");
    UnicodeString insertString("insert");
    baseString.insert( 2, insertString );
    EXPECT_EQ( baseString.toString(), "BAinsertSE");
  }
  catch(...)
  {
    FAIL() << "Did Not Expect Exception";
  }
}

TEST( UnicodeString, Append )
{
  try
  {
    UnicodeString baseString("地基");
    UnicodeString appendString("被加的字符串");
    baseString.append( appendString );
    EXPECT_EQ( baseString.toString(), "地基被加的字符串");
  }
  catch(...)
  {
    FAIL() << "Did Not Expect Exception";
  }
}

TEST( VarObject, BooleanConstructor )
{
  
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
