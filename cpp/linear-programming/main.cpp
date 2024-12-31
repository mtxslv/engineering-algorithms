#include <iostream>
#include "simplex.h"

using namespace std;

int main(){
    vector<vector<float>> T {
        {1.0f, -40.0f, -30.0f, 0.0f, 0.0f,  0.0f},
        {0.0f,   1.0f,   1.0f, 1.0f, 0.0f, 12.0f},
        {0.0f,   2.0f,   1.0f, 0.0f, 1.0f, 16.0f}
    };
    vector<vector<float>> Ttest {
        {1.0f,   0.0f,   -30.0f, 0.0f, 0.0f,  0.0f},
        {0.0f,   1.0f,    1.0f, 1.0f, 0.0f, 12.0f},
        {0.0f,   2.0f,    1.0f, 0.0f, 1.0f, 16.0f}
    };  
    
    int n = 2 ; // 2 variables
    int m = 2 ; // 2 constrains

    // int ansT = checkIfNegative(T,n);
    // int ansTtest = checkIfNegative(Ttest, n);
    // cout << "Tableau T has negative c values? " << ansT << endl;
    // cout << "Tableau Ttest has negative c values? " << ansTtest << endl ;
    printMatrix(T);
    simplexTableau(T,n,m);
    printMatrix(T);
}

/*
   1    -40    -30    0    0    0
   0      1      1    1    0   12
   0      2      1    0    1   16

*/