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
    for (int i = 0; i < outSize; ++i) {
        std::complex<double> sum(0.0, 0.0);
        for (int j = 0; j < m; ++j) {
            sum += A[i + j]*B[m - j - 1]; // Note the reversal of B
        }
        result[i] = sum;
    }

    return result;
}