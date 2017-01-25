#include <gtest/gtest.h>
#include "../include/StringUtils.hpp"

TEST( StringUtils, StripWhiteSpace ) {
  EXPECT_EQ("Spaces Removed From Beginning and End",StringUtils::stripWhiteSpace("  Spaces Removed From Beginning and End  "));
  EXPECT_EQ("",StringUtils::stripWhiteSpace("  ")); // String of Only Whitespace
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
