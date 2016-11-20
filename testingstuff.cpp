#include <iostream>

using namespace std;

int ztoi( string argString );

int main() {

}

int ztoi( string argString ) {
  int loopCount = 1;
  int runningTotal = 0;
  for ( int i = argString.length-2; i >= 0; i-=2 ) {
    if ( argString[i] == 0x4E && argString[i+1] == 0x00 ) {
      runningTotal = runningTotal + loopCount;
      loopCount = loopCount * 10;
      continue;
    }
    if ( argString[i] == 0x4E && argString[i+1] == 0x8C ) {
      runningTotal = runningTotal + 2*loopCount;
      loopCount = loopCount * 10;
      continue;
    }
    if ( argString[i] == 0x4E && argString[i+1] == 0x09 ) {
      runningTotal = runningTotal + 3*loopCount;
      loopCount = loopCount * 10;
    }
    if ( argString[i] == 0x56 && argString[i+1] == 0xDB ) {
      runningTotal = runningTotal + 4*loopCount;
      loopCount = loopCount * 10;
    }
    if ( argString[i] == 0x4E && argString[i+1] == 0x94 ) {
      runningTotal = runningTotal + 5*loopCount;
      loopCount = loopCount * 10;
    }
    if ( argString[i] == 0x51 && argString[i+1] == 0x6D ) {
      runningTotal = runningTotal + 6*loopCount;
      loopCount = loopCount * 10;
    }
    if ( argString[i] == 0x4E && argString[i+1] == 0x03 ) {
      runningTotal = runningTotal + 7*loopCount;
      loopCount = loopCount * 10;
    }
    if ( argString[i] == 0x51 && argString[i+1] == 0x6B ) {
      runningTotal = runningTotal + 8*loopCount;
      loopCount = loopCount * 10;
    }
    if ( argString[i] == 0x4E && argString[i+1] == 0x5D ) {
      runningTotal = runningTotal + 9*loopCount;
      loopCount = loopCount * 10;
    }
  }
}
