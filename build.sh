#!/bin/bash

RED='\033[0;31m';
GREEN='\033[0;32m';
NC='\033[0m';

install () {
  echo -n Building C Sources$'\t'$'\t'
  if gcc -c Sources/readFile.c -o Bin/readFile.o ; then
    echo -e [${GREEN}PASS${NC}]
  else
    echo -e [${RED}FAIL${NC}]
    echo Aborting Install...
    exit
  fi

  echo -n Building Swift Sources$'\t'$'\t'
  if swiftc -import-objc-header Sources/readFile.h Sources/main.swift Sources/varObject.swift Sources/mathParser.swift Sources/basicOutput.swift Sources/evaluateBoolean.swift Sources/stateMachine.swift Sources/storeVar.swift Sources/stripWhitespace.swift Bin/readFile.o -o Bin/cpl ; then
    echo -e [${GREEN}PASS${NC}]
  else
    echo -e [${RED}FAIL${NC}]
    echo Aborting Install...
    exit
  fi

  echo -n Installing to $1$'\t'$'\t'
  if sudo cp Bin/cpl $1 ; then
    echo -e [${GREEN}PASS${NC}]
  else
    echo -e [${RED}FAIL${NC}]
    echo Aborting Install...
    exit
  fi
}

uninstall () {
  echo Uninstalling from $1...
  if sudo rm $1 ; then
    echo Uninstall Successful
  else
    echo Uninstall Failed: Could not Remove $1
    exit
  fi
}

usage () {
  echo Usage: $./build.sh <install|uninstall> /path/to/installation
  echo Just enter ./build.sh for default installation.
  exit
}

if [ $# -eq 0 ]; then
  install /usr/bin/cpl
fi

if [ $# -eq 1 ]; then
  case $1 in
    install) install /usr/bin/cpl ;;
    uninstall) uninstall /usr/bin/cpl ;;
    *) usage ;;
  esac
fi

if [ $# -eq 2 ]; then
  case $1 in
    install) install $2 ;;
    uninstall) uninstall $2 ;;
    *) usage ;;
  esac
fi

if [ $# -ge 3 ]; then
  usage
fi
