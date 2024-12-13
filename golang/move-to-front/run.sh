#!/bin/bash

# Build the Go executable
go build main.go

# Initialize loop counters
x=1
y=1
trial=1

# Outer loop: Iterate over powers of 10 from 1 to 4
while [ $x -le 3 ]
do
  # Inner loop: Iterate from 1 to 10
  while [ $y -le 10 ]
  do
    # Calculate the size based on the current iteration
    size=$(( $y * 10**$x ))

    # Print the trial number and size
    echo "Trial #$trial: Size $size"

    # Execute the Go program with the specified arguments
    ./main equal $size $size>> experiments.txt

    # Increment the trial counter
    ((trial++))

    # Increment the inner loop counter
    ((y++))
  done

  # Reset the inner loop counter for the next outer loop iteration
  y=1

  # Increment the outer loop counter
  ((x++))
done


# Initialize loop counters
multiplier=1
NEWSIZE=900
while [ $multiplier -le 100 ]
do
    REQUESTS=$(( 180 * $multiplier ))
    # Execute the Go program with the specified arguments
    ./main diff $NEWSIZE $REQUESTS>> experiments.txt

    # Print the trial number and size
    echo "Trial #$trial: Size $NEWSIZE, Requests $REQUESTS"
    
    # Increment the trial counter
    ((trial++))
    
    ((multiplier++))

done