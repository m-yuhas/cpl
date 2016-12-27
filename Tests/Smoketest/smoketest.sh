#!/bin/bash

RED='\033[0;31m';
GREEN='\033[0;32m';
NC='\033[0m';

TEST_NAME_ARRAY=( "Printing" "If (False)" "If ( True )" "For 1" "For 2" "For 3" "For 4" "For 5" "While 0" "While 1" "While 2" "While 3" "While 4" "While 5" "While 6" "While 7" "While 8" "While 9" "Break" "Float Storage")
ANSWER_ARRAY=( "你好世界 4 4" "表达式否的" "表达式对的" "循环数：1" "循环数：2" "循环数：3" "循环数：4" "循环数：5" "0" "1" "2" "3" "4" "5" "6" "7" "8" "9" "101" "派=3.14159" )
OUTPUT="$(cpl smoketest)"

COUNT=0
while read -r LINE; do
  echo -n ${TEST_NAME_ARRAY[$COUNT]}$'\t'$'\t'$'\t'
  if [ "$LINE" == "${ANSWER_ARRAY[$COUNT]}" ] ; then
    echo -e [${GREEN}PASS${NC}]
  else
    echo -e [${RED}FAIL${NC}]
  fi
  COUNT=$((COUNT+1))
done <<< "$OUTPUT"
7
