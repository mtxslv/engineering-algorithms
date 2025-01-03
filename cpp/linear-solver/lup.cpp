#include <cmath>
#include <iomanip>
#include <iostream>
#include <vector>
#include "lup.h"
#include <math.h>

// Function to print a 2D vector (matrix)
void printMatrix(const std::vector<std::vector<float>>& matrix, int decimals) {
    for (const auto& row : matrix) {
        for (const auto& value : row) {
            if (decimals < 0) {
                std::cout << value << "\t";
            } else {
                if (value > -4e-06 && value < 5e-06) {
                    std::cout << static_cast<int>(round(value)) << "\t"; // simplify visualization
                }   else {
                    std::cout << std::setprecision(decimals) << value << "\t";
                }                 
            }
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
    std::vector<int>& p, // permutation array, row array
    std::vector<std::vector<float>>& b,  // col array
    std::vector<std::vector<float>>& x // result, col array
){
    
    // Define n 
    int n = L.size();

    // Define y 
    std::vector<std::vector<float>> y;
    y.resize(n, std::vector<float>(1,0.0f));
    x.resize(n, std::vector<float>(1,0.0f));

    float sum;

    for (int i=0; i<n; i++){
        sum = 0.0f;
        for (int j=0; j<i; j++){
            sum += L[i][j]*y[j][0];
        } 
        y[i][0] = b[p[i]][0] - sum;
    }

    for (int i = n-1; i >= 0; i--){
        sum = 0.0f;
        for (int j = i+1; j < n; j++) {
            sum += U[i][j]*x[j][0];
        }
        x[i][0] = (y[i][0]-sum)/U[i][i];
    }
}

void LupDecomposition(
    std::vector<std::vector<float>>& A,
    std::vector<int>& pi
) {

    int n = A.size();

    // Let \pi[1:n] be a new array
    pi.resize(n);

    for (int i = 0; i<n; i++){
        // initialize pi to the identity permutation
        pi[i] = i;
    }

    int k_;
    float p, aux;

    for (int k=0; k<n; k++){
        
        p = 0.0f;
        
        // find largest absolute value in column k
        for (int i=k; i<n; i++){
            if(abs(A[i][k]) > p) {
                p = abs(A[i][k]);
                // k_ is the row number of the largest found so far
                k_ = i ;
            }  
        }

        if (p == 0) {
            // raise error "singular matrix"
            throw std::invalid_argument("Singular Matrix");
        }

        // exchange pi[k] with pi[k_]
        aux = pi[k];
        pi[k] = pi[k_];
        pi[k_] = aux;

        // exchange rows k and k_
        for (int i=0; i<n; i++){
            aux = A[k][i];
            A[k][i] = A[k_][i];
            A[k_][i] = aux;
        }

        for (int i=k+1; i<n; i++) {
            A[i][k] = A[i][k] / A[k][k];

            // compute L and U in place in A
            
            for (int j = k+1; j<n; j++){
                A[i][j] = A[i][j] - A[i][k]*A[k][j];
            }
        
        }
    }   
}

void LupDecompAndTranscription(
    std::vector<std::vector<float>>& A,
    std::vector<std::vector<float>>& L, 
    std::vector<std::vector<float>>& U, 
    std::vector<int>& p // permutation array, row array   
) {
    // First, decompose A in-place
    LupDecomposition(A,p);

    // Now, let's transcribe A into L and U
    int n = A.size();

    for (int i=0 ; i<n; i++) {
        for (int j=0 ; j<n; j++) {
            // Upper matrix
            if (i>j){ 
                U[i][j] = 0;
            } else {
                U[i][j] = A[i][j];
            }
            // Lower matrix
            if (i<j) {
                L[i][j] = 0;
            } else {
                if (i == j) {
                    L[i][j] = 1 ;
                } else {
                    L[i][j] = A[i][j];
                }
            }
        }
    }
}


std::vector<std::vector<float>> LupInverse(
    std::vector<std::vector<float>> A
) {

    int n = A.size();

    // Inverse is also a nxn matrix
    std::vector<std::vector<float>> inverse(n,std::vector<float>(n));

    // LUP-decompose A
    std::vector<std::vector<float>> L(n,std::vector<float>(n));
    std::vector<std::vector<float>> U(n,std::vector<float>(n));
    std::vector<int> pi;

    LupDecompAndTranscription(A,L,U,pi);

    for(int i = 0; i < n; i++) {
        // Define c_i as the i-th column of a nxn identity matrix
        std::vector<std::vector<float>> c_i(n,std::vector<float>(1,0.0f));
        c_i[i][0] = 1;

        std::vector<std::vector<float>> x;

        lupSolver(L,U,pi,c_i,x);

        // Now we transcribe x into the inverse
        for (int j = 0; j < n; j++){
            inverse[j][i] = x[j][0];
        }
    }
    return inverse;
}