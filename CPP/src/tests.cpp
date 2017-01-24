#include <gtest/gtest.h>
#include "../include/StringUtils.hpp"

TEST( StripWhiteSpaceTest, General ) {
  EXPECT_EQ("thisstringhasnospace",StringUtils::stripWhiteSpace("this string has no space"));
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
