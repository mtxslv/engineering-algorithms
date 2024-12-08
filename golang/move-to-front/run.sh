#!/bin/bash

go build main.go

x=1
while [ $x -le 5 ]
do
  echo "Trial #$x"
  x=$(( $x + 1 ))
  ./main songs.csv >> experiments.txt
done