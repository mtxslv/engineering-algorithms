#!/bin/bash

SEEDA=1223334444
SEEDB=1224444333
N=2048
RESULTSFOLDER=./results
HOWMANYTHREADS='5 10 15'

if [ ! -d "$RESULTSFOLDER" ]; then 
    echo "$RESULTSFOLDER does not exist. Making it..."
    mkdir $RESULTSFOLDER
    mkdir $RESULTSFOLDER/parallel
    mkdir $RESULTSFOLDER/serial
fi

# COMPILE CODE
g++ -o generate generate.cpp
g++ -o matmultserial matmultserial.cpp
g++ -o matmultparallel -fopenmp matmultparallel.cpp

# GENERATE MATRICES
./generate $N $SEEDA ./matrix_a.txt
./generate $N $SEEDB ./matrix_b.txt

# RUN SERIAL PART
for i in {1..3}; do
    ./matmultserial ./matrix_a.txt ./matrix_b.txt $RESULTSFOLDER/serial/ans_$i.txt $RESULTSFOLDER/serial/log_$i.txt
done

# RUN PARALLEL PART
for numThread in $HOWMANYTHREADS; do
    export OMP_NUM_THREADS=$numThread
    echo "Using $numThread threads."
    resultsFolderPath=$RESULTSFOLDER/parallel/$numThread
    if [ ! -d "$resultsFolderPath" ]; then 
        echo "$resultsFolderPath does not exist. Making it..."
        mkdir $resultsFolderPath
    fi
    for i in {1..3}; do
        ./matmultparallel ./matrix_a.txt ./matrix_b.txt $resultsFolderPath/ans_$i.txt $resultsFolderPath/log_$i.txt
    done
done


#https://stackoverflow.com/questions/59838/how-do-i-check-if-a-directory-exists-or-not-in-a-bash-shell-script