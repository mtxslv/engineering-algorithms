#include <iostream>
#include <vector>
#include "lup.h"

// Function to print a 2D vector (matrix)
void printMatrix(const std::vector<std::vector<float>>& matrix) {
    for (const auto& row : matrix) {
        for (const auto& value : row) {
            std::cout << value << " ";
        }
        std::cout << std::endl;
    }
}

void lupSolver(
    std::vector<std::vector<float>>& L, 
    std::vector<std::vector<float>>& U, 
    std::vector<std::vector<float>>& p, // permutation array
    std::vector<std::vector<float>>& b, 
    std::vector<std::vector<float>>& x, // result
    int n
){
    std::cout <<"OK" <<std::endl;
}