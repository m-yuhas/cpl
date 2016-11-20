#include <string>

using namespace std;

class Var {
  public:
    char type;
    float floatVal;
    int intVal;
    bool boolVal;
    string stringVal;
    Var(float initialValue);
    Var(int initialValue);
    Var(bool initialValue);
    Var(string initialValue);
    Var();
    void updateValue( float updateVal );
    void updateValue( int updateVal );
    void updateValue( bool updateVal );
    void updateValue( string updateVal );
    char getType();
};

Var::Var(float initialValue) {
  type = 3;
  floatVal = initialValue;
  return;
}

Var::Var(int initialValue) {
  type = 2;
  intVal = initialValue;
  return;
}

Var::Var(bool initialValue) {
  type = 1;
  boolVal = initialValue;
  return;
}

Var::Var(string initialValue) {
  type = 4;
  stringVal = initialValue
  return;
}

Var::Var() {
  type = 0;
  return;
}

void Var::updateValue(float updateVal) {
  type = 3;
  floatVal = updateVal;
  return;
}

void Var::updateValue(int updateVal) {
  type = 2;
  intVal = updateVal;
  return;
}

void Var::updateValue(bool updateVal) {
  type = 1;
  boolVal = updateVal;
  return;
}

void Var::updateValue(string updateVal) {
  type = 4;
  stringVal = updateVal;
  return;
}

char Var::getType() {
  return type;
}
