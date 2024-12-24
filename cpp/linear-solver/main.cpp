#include <iostream>
#include <vector>
#include "file1.h"
#include "lup.h"

using namespace std;

void demo(){
    std::cout << "This is a demo" << std::endl;
    danDaDan();
    std::cout << std::endl;
    onePiece();
    std::cout << std::endl;
}


void testingLUPsolver(){
    vector<vector<float>> L;
    vector<vector<float>> U;
    // Create and initialize the 2D vector (matrix)
    vector<vector<float>> matrix = {
        {1.0f, 0.0f, 0.0f},
        {0.2f, 2.0f, 0.0f},
        {0.6f, 0.5f, 1.0f}
    };

    // Print the matrix
    printMatrix(matrix);
}

int main(){
    demo();
    testingLUPsolver();
    return 0;
}