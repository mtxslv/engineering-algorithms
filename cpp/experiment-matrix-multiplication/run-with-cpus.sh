#!/bin/bash
#SBATCH --partition=amd-512  # partição para a qual o job é enviado
#SBATCH --time=0-4:0    # Especifica o tempo máximo de execução do job, dado no padrão dias-horas:minutos
#SBATCH --cpus-per-task=5

SEEDA=1223334444
SEEDB=1224444333
N=4
RESULTSFOLDER=./results

if [ ! -d "$RESULTSFOLDER" ]; then 
    echo "$RESULTSFOLDER does not exist. Making it..."
    mkdir $RESULTSFOLDER
    mkdir $RESULTSFOLDER/parallel
fi

# COMPILE CODE
g++ -o generate generate.cpp
g++ -o matmultparallel -fopenmp matmultparallel.cpp

# GENERATE MATRICES
./generate $N $SEEDA ./matrix_a.txt
./generate $N $SEEDB ./matrix_b.txt

# RUN PARALLEL PART
export OMP_NUM_THREADS=$SLURM_CPUS_PER_TASK
echo "Using $OMP_NUM_THREADS threads."
resultsFolderPath=$RESULTSFOLDER/parallel/$numThread
if [ ! -d "$resultsFolderPath" ]; then 
    echo "$resultsFolderPath does not exist. Making it..."
    mkdir $resultsFolderPath
fi
for i in {1..3}; do
    ./matmultparallel ./matrix_a.txt ./matrix_b.txt $resultsFolderPath/ans_$i.txt $resultsFolderPath/log_$i.txt
done

#https://stackoverflow.com/questions/59838/how-do-i-check-if-a-directory-exists-or-not-in-a-bash-shell-script