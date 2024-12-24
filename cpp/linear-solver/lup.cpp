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

std::vector<int> permutationToArray(std::vector<std::vector<float>>& P){
    /*
    For i = 1,2,3, ..., n, the entry p[i] indicates that P_{i,p[i]} = 1
    and P_{ij} = 0 for j \neq p[i]
    */
   int n = P.size();
   std::vector<int> p;
   p.resize(1,n);
   std::cout << "PERMUTATION MATRIX HAS SIZE: " << n<<"x"<<P[0].size() <<std::endl;
   for (int i = 0; i < n; i++) {
     for (int j = 0; j < n; j++) {
        if (P[i][j] == 1) {
            p[i] = j;
            break;
        }
     }
   }
   return p;
}

void lupSolver(
    std::vector<std::vector<float>>& L, 
    std::vector<std::vector<float>>& U, 
    std::vector<int>& p, // permutation array
    std::vector<std::vector<float>>& b, 
    std::vector<std::vector<float>>& x, // result
    int n
){
    std::cout <<"OK" <<std::endl;
}