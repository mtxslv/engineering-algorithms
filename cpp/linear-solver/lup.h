#ifndef LUP_H
#define LUP_H

void lupSolver(
    std::vector<std::vector<float>>& L, 
    std::vector<std::vector<float>>& U, 
    std::vector<int>& p, // permutation array
    std::vector<std::vector<float>>& b, 
    std::vector<std::vector<float>>& x, // result
    int n
);

void printMatrix(const std::vector<std::vector<float>>& matrix);

std::vector<int> permutationToArray(std::vector<std::vector<float>>& P);

#endif