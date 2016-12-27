#!/bin/bash

N_PREV=1
FIB_NUM=1
TEMP=0
COUNTER=3

while [ $COUNTER -lt 93 ]; do
  TEMP=$((FIB_NUM))
  FIB_NUM=$((N_PREV+FIB_NUM))
  N_PREV=$((TEMP))
  COUNTER=$((COUNTER+1))
done

echo $FIB_NUM
