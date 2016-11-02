#include <string>

using namespace std;

class Utils {
  public:
    static string stripWhiteSpace( string argString );
    static string print( string argString );
};

string Utils::stripWhiteSpace( string argString ) {
  string returnString = "";
  bool inQuotes = false;
  for ( int i = 0; i < argString.length(); i++ ) {
    if ( inQuotes == true ) {
      //returnString.append(&argString[i]);
      returnString += argString[i];
      if ( argString[i] == '"' ) {
        inQuotes = false;
      }
      continue;
    }
    if ( inQuotes == false ) {
      if ( argString[i] != ' ' && argString[i] != '\t' ) {
        //returnString.append(&argString[i]);
        returnString += argString[i];
      }
      if ( argString[i] == '"' ) {
        inQuotes = true;
      }
      continue;
    }
  }
  return returnString;
}

string Utils::print( string argString ) {
  string returnString = "";
  bool inQuotes = false;
  for ( int i=0; i < argString.length(); i++ ) {
    if ( inQuotes == true ) {
      if ( argString[i] == '"' ) {
        inQuotes = false;
        //returnString += "Also";
      } else {
        returnString += argString[i];
        //returnString += "A";
      }
      continue;
    }
    if ( inQuotes == false ) {
      if ( argString[i] == '"' ) {
        inQuotes = true;
        //returnString += "HERE";
        //cout << i << '\n';
      } else {
        return "\033[1;31mERROR: Not Yet Supported\033[0m";
        //return returnString;
      }
      continue;
    }
  }
  return returnString;
}
