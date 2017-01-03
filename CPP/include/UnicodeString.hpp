// Test Class to Implement Unicode Strings in CPP
// (C) 2016 Michael Yuhas

#include <vector>
#include <string>
#include "../include/InvalidUTF8Exception.hpp"

/*
UnicodeString: Method For Storing a Unicode String
Public Methods:
Constructor - constructs a UnicodeString Object
append - appends one UnicodeString Object to Another
substring - returns a UnicodeString which is the substring between two indeces
getChar - returns a string of the unicode character at an index
writeChar - overwrites the unicode character at an index
removeChar - removes the unicode character at an index
toString - writes the UnicodeString object as a regular string
Protected Attributes:
uniString - a vector of strings, each representing a unicode character
*/
class UnicodeString
{
  public:
    UnicodeString( std::string inputString );
    void append( UnicodeString stringToAppend );
    void substring( UnicodeString substring );
    void getChar( int index );
    void writeChar( int index );
    void removeChar( int index );
    void insert( int insertLocation, UnicodeString stringToInsert );
    std::string toString();

  protected:
    std::vector<std::string> uniString;
    std::vector<std::string> getVector();
};
