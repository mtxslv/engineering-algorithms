#include <cfloat> // To use FLT MAX
#include <iostream>
#include <vector>
#include "simplex.h"

using namespace std;

bool checkIfNegative(vector<vector<float>>& T, int n){
    // Return negative if any value of T[0,1:n] is still less than 0
    for (int i = 1; i < n+1 ; i++) { // n variables. T[0][0] is Z
        if (T[0][i] < 0) {
            return true;
        }
    }
    return false;
}


void simplexTableau(
    vector<vector<float>>& T,
    int n,
    int m
) {
    // Quotient
    vector<float> q;
    q.resize(m);
    float qCompare = FLT_MAX;

    int iP = -1; // Row where smallest quotient happens
    
    // Variables to find col where largest c_i happens
    int jP = -1;
    float jPcompare = FLT_MAX;

    // while (checkIfNegative(T,n)) {
        // First, find in which column the largest c_i happens
        for (int i = 1; i < n+1 ; i++) { // n variables. T[0][0] is Z
            if (T[0][i] < jPcompare) {
                jPcompare = T[0][i];
                jP = i;
            }    
        }
        // Let's populate the quotient vector
        for (int i=1; i<m+1; i++) {
            if (T[i][jP] <= 0) {
                q[i-1] = FLT_MAX;
            } else {
                // element wise division between the restriction rows
                // on the pivot column and the last tableau column
                q[i-1] = T[i][jP] / T[i][m+n+1]; 
            }
        }

        // Now we found the smallest restriction quotient
        for (int i=0; i<m; i++){
            if (q[i] < qCompare){
                iP = i+1; // q[i] is shiftted backwards, hence the +1
                qCompare = q[i];
            }
        }

        cout << "The Pivot element is T["<<iP<<"]["<<jP<<"] = " << T[iP][jP] << endl;

        // Now we pivot 
        for (int i=0; i<m+1;i++){
            for (int j=0; j<n+m+2; j++){
                if (i == iP) {
                    T[i][j] = T[i][j]/T[iP][jP];
                } else {
                    T[i][j] = T[i][j] - (T[i][jP]/T[iP][jP])*T[iP][j];
                }
            }
        }
    // }
}

// Function to print a 2D vector (matrix)
void printMatrix(const std::vector<std::vector<float>>& matrix) {
    for (const auto& row : matrix) {
        for (const auto& value : row) {
            std::cout << value << "\t";
        }
        std::cout << std::endl;
    }
}