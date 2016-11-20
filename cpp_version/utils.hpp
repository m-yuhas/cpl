#include <string>
#include "var.hpp"

using namespace std;

class Utils {
  public:
    static string stripWhiteSpace( string argString );
    static string print( string argString, std::map<string, Var> &varMap );
    static string store( string argString, std::map<string, Var> &varMap );
    static Var evaluateExpression( string expression, std::map<string, Var> &varMap );
    static Var chineseNumberToVar( string argString );
    static Var westernNumberToVar( string argString );
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

string Utils::print( string argString, std::map<string, Var> &varMap ) {
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

string Utils::store( string argString, std::map<string, Var> &varMap) {
  for ( int i=0; i < argString.length(); i++ ) {
    if ( argString == '"' || argString == "'" || argString == "+" || argString == "-" || argString == "/" || argString == "\\" || argString == "[" || argString == "]" || argString == "(" || argString== ")" || argString == "{" || argString == "}" || argString == "&" || argString == "!" || argSring == "<" || argString == ">" || argString == "." || argString == "," ) {
      return "\033[1,31m变量名称不合适\033[0m"
    }
  }
  string inVal;
  cin >> inVal;
  if //TODO Implement Test
  varMap[argString] = inVal;
  return "";
}

Var Utils::evaluateExpression( string expression, std::map<string, Var> &varMap ) {
  return NULL;
}

Var Utils::stringToNumber( string expression, std::map<string, Var> &varMap ) {
  int bins[3];
  int numberOfDecimals;
  for ( int i=0; i < expression.length(); i++ ) {
    if ( expression[i] == '1' || expression[i] == '2' || expression[i] == '3' || expression[i] == '4' || expression[i] == '5' || expression[i] == '6' || expression[i] == '7' || expression[i] == '8' || expression[i] == '9' || expression == '0' ) {
      bins[0]++;
      continue;
    }
    if ( i < expression.length()-1 ) {
      if ( expression[i] == 0x4E && ( expression[i+1] == 0x00 || expression[i+1] == 0x8C || expression[i+1] == 0x09 || expression[i+1] == 0x94 || expression[i+1] == 0x03 || expression[i+1] == 0x5D ) ) {
        bins[1]++;
        i++;
      }
      if ( expression[i] == 0x56 && ( expression[i+1] == 0xDB ) ) {
        bins[i]++;
        i++;
      }
      if ( expression[i] == 0x51 && ( expression[i+1] == 0x6D || expression[i+1] == 0x6B ) ) {
        bins[i]++;
        i++;
      }
      if ( expression[i] == 0x96 && expression[i+1] == 0xF6 ) {
        bins[i]++;
        i++;
      }
      if ( expression[i] == 0x30 && expression[i+1] == 0x02 ) {
        numberOfDecimals++;
        i++;
      }
      continue;
    }
    if ( expression[i] == '.' ) {
      numberOfDecimals++;
      continue;
    }
    bins[2]++;
  }
  if ( bins[0] > 0 && bins[1] == 0 && bins[2] == 0 ) {
    if ( numberOfDecimals == 0 ) {
      for ( int j = expression.length(); j > 0; j-=2 ) {
        
      }
    }
  }
  if ( bins[0] == 0 && bins[1] > 0 && bins[2] == 0 ) {

  }

  return;
}
