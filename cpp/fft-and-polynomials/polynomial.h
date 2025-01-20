#ifndef POLYNOMIAL_H
#define POLYNOMIAL_H

#include <vector>

using namespace std;

double poly(double x, vector<double> params);
std::vector<std::complex<double>> conv1D(
    std::vector<std::complex<double>> A,
    std::vector<std::complex<double>> B
);
vector<double> polynomialProduct ( vector<double> A, vector<double> B);

#endif