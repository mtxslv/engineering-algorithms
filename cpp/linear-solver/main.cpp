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
    // Create and initialize the 2D vector (matrix)
    vector<vector<float>> L = {
        {1.0f, 0.0f, 0.0f},
        {0.2f, 1.0f, 0.0f},
        {0.6f, 0.5f, 1.0f}
    };
    vector<vector<float>> U = {
        {5.0f, 6.0f,  3.0f},
        {0.0f, 0.8f, -0.6f},
        {0.0f, 0.0f,  2.5f},
    };
    vector<vector<float>> P = {
        {0.0f, 0.0f, 1.0f},
        {1.0f, 0.0f, 0.0f},
        {0.0f, 1.0f, 0.0f},
    };
    vector<vector<float>> b = {
        {3},
        {7},
        {8},
    };
    vector<vector<float>> x;
    vector<int> p = permutationToArray(P); // ANS MUST BE [2, 0, 1]
    lupSolver(L,U,p,b,x);
    for (int k = 0; k < 3 ; k++){
        std::cout << "ANS = " << x[k][0] << std::endl;
    }
    // Print the matrix
    printMatrix(P);    
    std::cout << "L U P DECOMPOSITION" << std::endl;
    vector<vector<float>> A = {
        {1,2,0},
        {3,4,4},
        {5,6,3},
    };
    // LupDecomposition(A);
    vector<vector<float>> Lans, Uans;
    vector<int> pi;
    LupDecompAndTranscription(A,L,U,pi);
    std::cout << "======== L DECOMP =======" << std::endl;
    printMatrix(L);
    std::cout << "======== U DECOMP =======" << std::endl;
    printMatrix(U);    
    std::cout << "======== A =======" << std::endl;
    printMatrix(A);
    std::cout << "======== PI =======\n[\n" ;
    for (int k = 0; k < 3 ; k++){
        std::cout << pi[k] << ",";
    }
    std::cout <<"\n]\n";
}

int main(){
    demo();
    testingLUPsolver();
    return 0;
}