#include <gtest/gtest.h>
#include "../include/StringUtils.hpp"

TEST( StringUtils, StripWhiteSpaceFromBeginningAndEnd )
{
  EXPECT_EQ("Spaces Removed From Beginning and End",StringUtils::stripWhiteSpace("  Spaces Removed From Beginning and End  "));
}

TEST( StringUtils, StripWhiteSpaceFromAllWhitespaceString )
{
  EXPECT_EQ("",StringUtils::stripWhiteSpace("  "));
}
