// Invalid UTF-8 String Exception
// (C) 2016 Michael Yuhas
#include <string>

using namespace std;

class InvalidUTF8Exception
{
  public:
    InvalidUTF8Exception( string badString );
    string what();
  protected:
    string badString;
}

InvalidUTF8Exception::InvalidUTF8Exception( string badString ) {
  this.badString = badString;
}

InvalidUTF8Exception::what() {
  return "Invalid UTF-8 String: \"" + this.badString + "\"";
}
