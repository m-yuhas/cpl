#!/bin/bash

RED='\033[0;31m';
GREEN='\033[0;32m';
NC='\033[0m';
PASSES=0;

echo -n Building C Sources$'\t'$'\t'
if gcc readFile.c -c ; then
  echo -e [${GREEN}PASS${NC}]
  PASSES+=1;
else
  echo -e [${RED}FAIL${NC}]
  echo Aborting Install...
  exit;
fi

echo -n Building Swift Sources$'\t'$'\t'
if swiftc -import-objc-header readFile.h main.swift varObject.swift mathParser.swift basicOutput.swift evaluateBoolean.swift stateMachine.swift storeVar.swift readFile.o -o cpl ; then
  echo -e [${GREEN}PASS${NC}]
  PASSES+=1;
else
  echo -e [${RED}FAIL${NC}]
  echo Aborting Install...
  exit;
fi

echo -n Installing to /usr/bin$'\t'$'\t'
if sudo cp cpl /usr/bin/cpl ; then
  echo -e [${GREEN}PASS${NC}]
else
  echo -e [${RED}FAIL${NC}]
  echo Aborting Install...
  exit;
fi
