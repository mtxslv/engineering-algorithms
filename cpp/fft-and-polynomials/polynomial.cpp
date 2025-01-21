#include <algorithm>
#include <math.h>       /* pow */
#include <complex>
#include <vector>
#include <iostream>

using namespace std;

double poly(double x, vector<double> params){
    double y = 0.0;
    int n = params.size();
    for (int k = 0; k < n; k++) {
        y += params[k]*pow(x,k);
    }
    return y;
}

std::vector<std::complex<double>> conv1D(
    std::vector<std::complex<double>> A,
    std::vector<std::complex<double>> B
) {
    // Ensure A is the larger vector (if not, swap them)
    if (A.size() < B.size()) {
        std::swap(A,B);
    }

    int n = A.size();
    int m = B.size();
    int outSize = n + m - 1; // Size of the output vector for 'valid' convolution
    std::vector<std::complex<double>> result(outSize);

    // Perform the convolution
    int indexA, indexB;
    for (int k = 0; k < outSize; k++) {
        std::complex<double> sum(0.0, 0.0);
        for (int i = 0 ; i < n; i++) {
            indexA = i ;
            indexB = k-i; // Note the reversal of B
            if (indexA >= 0 && indexA < n && indexB >= 0 && indexB < m) {
                // If the index are valid (inside the limits) we consider
                sum += A[indexA]*B[indexB]; 
            } 
        }
        result[k] = sum;
    }
    return result;
}

std::vector<double> conv1D(
    std::vector<double> A,
    std::vector<double> B
) {
    // Ensure A is the larger vector (if not, swap them)
    if (A.size() < B.size()) {
        std::swap(A,B);
    }

    int n = A.size();
    int m = B.size();
    int outSize = n + m - 1; // Size of the output vector for 'valid' convolution
    std::vector<double> result(outSize);

    // Perform the convolution
    int indexA, indexB;
    for (int k = 0; k < outSize; k++) {
        double sum = 0.0;
        for (int i = 0 ; i < n; i++) {
            indexA = i ;
            indexB = k-i; // Note the reversal of B
            if (indexA >= 0 && indexA < n && indexB >= 0 && indexB < m) {
                // If the index are valid (inside the limits) we consider
                sum += A[indexA]*B[indexB]; 
            } 
        }
        result[k] = sum;
    }
    return result;
}

vector<double> polynomialProduct ( vector<double> A, vector<double> B) {
    int n = A.size()-1;
    int m = B.size()-1;
    int outSize = n + m + 1;

    vector<double> C(outSize);

    for(int i = 0 ; i <= n ; i++) {
        for (int j = 0; j <= m; j++) {
            C[i+j] += A[i]*B[j];
        }
    }
    return C;
}

void printPoly(char icon, vector<double> P){
    cout << icon << "(x) = " << P[0] ;

    for(int i = 1; i < P.size(); i++){
        if (P[i] >= 0) {
            cout << " + " ;
        } else {
            cout << " ";
        }
        cout << P[i] << "*x^" << i;
    }
    cout << endl;
}

vector<complex<double>> pointwiseProd(
    vector<complex<double>> A,
    vector<complex<double>> B
) {
    // Ensure A is the larger vector (if not, swap them)
    if (A.size() < B.size()) {
        std::swap(A,B);

        // Now pad B
        for (int k=0; k < A.size() - B.size(); k++){
            B.push_back(complex<double>(0.0));
        }
    }

    int n = A.size();
    vector<complex<double>> C(n);

    for (int k=0; k < n; k++){
        C[k] = A[k]*B[k];
    }

    return C;
}