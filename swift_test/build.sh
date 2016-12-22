#!/bin/bash
echo Building C Sources
if gcc readFile.c -c ; then
  echo [PASS]
else
  echo [FAIL]
fi

echo Building Swift Sources
if swiftc -import-objc-header readFile.h main.swift varObject.swift mathParser.swift readFile.o -o cpl ; then
  echo [PASS]
else
  echo [FAIL]
fi
