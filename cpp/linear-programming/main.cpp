#include <iostream>
#include "simplex.h"

using namespace std;

int main(){
    vector<vector<float>> T {
        {1.0f, -4.0f, -3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f},
        {0.0f,  2.0f,  3.0f, 1.0f, 0.0f, 0.0f, 0.0f, 6.0f},
        {0.0f, -3.0f,  2.0f, 0.0f, 1.0f, 0.0f, 0.0f, 3.0f},
        {0.0f,  0.0f,  2.0f, 0.0f, 0.0f, 1.0f, 0.0f, 5.0f},
        {0.0f,  2.0f,  1.0f, 0.0f, 0.0f, 0.0f, 1.0f, 4.0f}
    }; 
    simplexTableau(T,2,4); // 2 variables and 4 constraints
    printMatrix(T);
    cout << endl << endl << "FINAL MATRIX::" << endl;
    printMatrix(T);
}

/*
   1    -40    -30    0    0    0
   0      1      1    1    0   12
   0      2      1    0    1   16

    vector<vector<float>> T {
        {1.0f, -40.0f, -30.0f, 0.0f, 0.0f,  0.0f},
        {0.0f,   1.0f,   1.0f, 1.0f, 0.0f, 12.0f},
        {0.0f,   2.0f,   1.0f, 0.0f, 1.0f, 16.0f}
    }; 
    simplexTableau(T,2,2); // 2 variables and 2 constraints
    printMatrix(T);
    cout << endl << endl << "FINAL MATRIX::" << endl;
    printMatrix(T);
*/