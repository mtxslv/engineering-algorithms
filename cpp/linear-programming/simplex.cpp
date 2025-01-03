#include <cfloat> // To use FLT MAX
#include <iostream>
#include <iomanip>
#include <string>
#include <vector>
#include <sstream>
#include "simplex.h"

using namespace std;

bool checkIfNegative(vector<vector<float>>& T, int n, int m){
    // Return negative if any value of T[0,1:n+m+1] is still less than 0
    // That is, if any variable's multiplier is negative
    for (int i = 1; i < n+m+1 ; i++) { // n variables. T[0][0] is Z
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
    float jPcompare;

    // Pivot auxiliary variables;
    float pivot, multiplier ;

    while (checkIfNegative(T,n,m)) {
        // Reset comparison values
        jPcompare = FLT_MAX;
        qCompare = FLT_MAX;

        // First, find in which column the largest c_i happens
        for (int i = 1; i < n+m+1 ; i++) { // Here we consider the slack variables too. Moreover, T[0][0] is Z
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
                q[i-1] = T[i][m+n+1]/ T[i][jP]; 
            }
        }

        // Now we found the smallest restriction quotient
        for (int i=0; i<m; i++){
            if (q[i] < qCompare){
                iP = i+1; // q[i] is shiftted backwards, hence the +1
                qCompare = q[i];
            }
        }

        // Now we pivot 
        pivot = T[iP][jP];

        // First, make pivot value equal to 1
        for (int j= 0; j < n+m+2 ; j++) {
            T[iP][j] = T[iP][j]/pivot;
        }

        // Use the pivot row to pivot the rest of the matrix
        for (int i=0; i<m+1;i++){
            multiplier = T[i][jP];
            for (int j=0; j<n+m+2; j++){
                if (i != iP) {
                    T[i][j] = T[i][j] - multiplier*T[iP][j];
                }
            }
        }
    }
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

std::string processOutput(std::vector<std::vector<float>>& T, int n, int m) {
    std::string ans = "Maximized objective value: ";
    
    std::stringstream ss;
    ss << ans << std::fixed << std::setprecision(5) << T[0][m + n + 1] << endl;
    

    vector<float> sum;
    sum.resize(n+m); // input and slack variables
    vector<float> pos; // positions of the values on the last col
    pos.resize(n+m);

    // Let's check which variables are being used
    // A variable used must have a single 1 on the column
    for (int i = 0; i < m+1 ; i++) {
        for (int j = 0; j < n+m; j++) {
            sum[j] += T[i][j+1]; // j+1 because the 0-th col is var z
            // If it is 1, keep row number
            if (T[i][j+1] == 1) {
                pos[j] = i;
            }
        }
    } 
    // Now we look for the desired values
    int inputVarI = 1;
    int slackVarI = 1;
    for (int j = 0; j < n+m; j++) {
        if(j < n) { // we are dealing with the input variables
            ss << "x" << inputVarI << " = " ;
            if(sum[j] == 1){
                ss << std::setprecision(5) << T[pos[j]][m + n + 1];
            } else {
                ss << std::setprecision(5) << 0.0;
            }
            ss << endl;
            inputVarI++;
        } else { // we are dealing with the slack variables
            ss << "s" << slackVarI << " = " ;
            if(sum[j] == 1){
                ss << std::setprecision(5) << T[pos[j]][m + n + 1];
            } else {
                ss << std::setprecision(5) << 0.0;
            }
            ss << endl;
            slackVarI++;
        }
    }
    
    return ss.str();
}