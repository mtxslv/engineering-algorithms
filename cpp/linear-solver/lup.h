#ifndef LUP_H
#define LUP_H

void lupSolver(
    std::vector<std::vector<float>>& L, 
    std::vector<std::vector<float>>& U, 
    std::vector<std::vector<float>>& p, // permutation array
    std::vector<std::vector<float>>& b, 
    std::vector<std::vector<float>>& x, // result
    int n
);

void printMatrix(const std::vector<std::vector<float>>& matrix);

#endif