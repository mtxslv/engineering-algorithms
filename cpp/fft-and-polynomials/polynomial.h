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
std::vector<double> conv1D(
    std::vector<double> A,
    std::vector<double> B
);
void printPoly(char icon, vector<double> P);

#endif