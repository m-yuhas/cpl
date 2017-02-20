// String Utilities
// (C) 2016 Michael Yuhas

#include <string>

/*
StringUtils: Utilities to make string processing easier
Public Methods:
stripWhiteSpace - removes whitespace from the ends of a string
removeWhiteSpaceNotInQuotes - remove whitespace not contained within quotation marks
*/
class StringUtils
{
  public:
    static std::string stripWhiteSpace( std::string inputString );
    static std::string removeWhiteSpaceNotInQuotes( std::string inputString );
    static std::string stod( std::string inputString );
};
