#!/bin/bash

SEEDA=1223334444
SEEDB=1224444333
N=2048

g++ -o generate generate.cpp
g++ -o matmult matmultserial.cpp

./generate $N $SEEDA ./matrix_a.txt
./generate $N $SEEDB ./matrix_b.txt
./matmult ./matrix_a.txt ./matrix_b.txt ./ans.txt