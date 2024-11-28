#!/bin/bash

SEEDA=1223334444
SEEDB=1224444333
N=4
RESULTSFOLDER=./results

if [ ! -d "$RESULTSFOLDER" ]; then 
    echo "$RESULTSFOLDER does not exist. Making it..."
    mkdir $RESULTSFOLDER
    mkdir $RESULTSFOLDER/parallel
    mkdir $RESULTSFOLDER/serial
fi

# COMPILE CODE
g++ -o generate generate.cpp
g++ -o matmult matmultserial.cpp

# GENERATE MATRICES
./generate $N $SEEDA ./matrix_a.txt
./generate $N $SEEDB ./matrix_b.txt

# RUN SERIAL PART
for i in {1..3}; do
    ./matmult ./matrix_a.txt ./matrix_b.txt $RESULTSFOLDER/serial/ans_$i.txt
done

#https://stackoverflow.com/questions/59838/how-do-i-check-if-a-directory-exists-or-not-in-a-bash-shell-script