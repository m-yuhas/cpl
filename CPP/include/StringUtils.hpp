// String Utilities
// (C) 2016 Michael Yuhas

#include <string>

class StringUtils
{
  public:
    static std::string stripWhiteSpace( std::string inputString );
    static std::string removeWhiteSpaceNotInQuotes( std::string inputString );
};
