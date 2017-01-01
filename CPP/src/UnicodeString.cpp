// Test Class to Implement Unicode Strings in CPP
// (C) 2016 Michael Yuhas
#include <vector>
#include <string>
#include "../include/InvalidUTF8Exception.hpp"

using namespace std;

class UnicodeString
{
  public:
    UnicodeString( string inputString );
    void append( UnicodeString stringToAppend );
    void substring( UnicodeString substring );
    void getChar( Int index );
    void writeChar( Int index );
    void removeChar( Int index );

  protected:
    std::vector<string> uniString;
};

UnicodeString::UnicodeString( string inputString ) {
  for ( int i=0; i<inputString.size(); i++ ) {
   if ( (inputSring[i] | 0x1F) & 0xE0 == 0xC0 ) {
     if ( i+1 < inputString.size() && (inputString[i+1] | 0x3F) & 0xC0 == 0x80) {
        string charString ( inputString[i], inputString[i+1] );
        uniString.push_back( charString );
        i++;
     } else {
       throw InvalidUTF8Exception( inputString );
     }
   } else if ( (inputString[i] | 0x0F) & 0xF0 == 0xE0 ) {

   } else if ( (inputString[i] | 0x07) & 0xF8 == 0xF0 ) {

   } else if ( (inputString[i] | 0x7F) & 0x80 == 0x00 ) {
     string charString ( inputString[i] )
     uniString.push_back( charString );
   } else {
     throw InvalidUTF8Exception( inputString );
   }
  }
}
