#ifndef FFT_H
#define FFT_H

#include <complex>
#include <vector>

using namespace std;

vector<complex<double>> rootsOfUnity(int n);
vector<complex<double>> DFT(vector<double> a);
vector<complex<double>> FFT(vector<double> a);
vector<complex<double>> safeFFT(vector<double> a);
int nextPowerOf2(int n);
vector<double> padToPowerOf2(vector<double> a);

#endif