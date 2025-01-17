#ifndef FFT_H
#define FFT_H

#include <complex>
#include <vector>

using namespace std;

vector<complex<double>> DFT(vector<double> a);
vector<complex<double>> FFT(vector<double> a);

#endif