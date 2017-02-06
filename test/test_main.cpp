#include <gtest/gtest.h>
#include "test_StringUtils.cpp"
#include "test_UnicodeString.cpp"
#include "test_VarObject.cpp"
//#include "../include/VarObject.hpp"
//#include "../include/InvalidOperationException.hpp"

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();

  
}
