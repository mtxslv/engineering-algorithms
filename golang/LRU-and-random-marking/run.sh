#!/bin/bash

# Build the Go executable
go build main.go

# Initialize loop counter
trial=1

while [ $trial -le 10 ]
do
    # Execute the go program with args
    ./main ../matrix-mult/mem-access-matmult.txt >> results.txt
    
    # Increment the trial counter
    ((trial++))
done
