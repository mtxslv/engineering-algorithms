#include <iomanip>
#include <iostream>
#include <cstring>
#include <vector>
#include "fft.h"
#include "polynomial.h"
#include "taylorSeries.h"

using namespace std;

int main(int argc, char* argv[]) {
    vector<double> a = {1.0, 2.0, 3.0, 4.0, 5.0};
    vector<complex<double>> dftA = DFT(a);
    for (int k = 0; k < 4; k++) {
        cout << dftA[k] << endl;
    }
}