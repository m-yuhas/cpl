#!/usr/bin/python

n_prev=1
fib_num=1
temp=0

for i in range(3,93):
    temp=fib_num
    fib_num=n_prev+fib_num
    n_prev=temp

print(str(fib_num))
