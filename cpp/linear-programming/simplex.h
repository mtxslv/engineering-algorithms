#ifndef SIMPLEX_H
#define SIMPLEX_H

#include <vector>

void simplexTableau(
    std::vector<std::vector<float>>& T,
    int n,
    int m
);

void firstFeasibleSolutionMinimization(
    std::vector<std::vector<float>>& T,
    int n,
    int m    
);

bool checkIfNegative(std::vector<std::vector<float>>& T, int n, int m);

void printMatrix(const std::vector<std::vector<float>>& matrix);

std::string processOutput(std::vector<std::vector<float>>& T, int n, int m);

#endif