#include <chrono>
#include <cmath>
#include <cstring>
#include <iomanip>
#include <iostream>
#include <vector>
#include "fft.h"
#include "polynomial.h"
#include "taylorSeries.h"

using namespace chrono;
using namespace std;

void taylorSeriesEx() { 
    int howManyTerms = 32;

    vector<double> params = naturalLogarithm(howManyTerms);
    double y = poly(0.3, params); // 0.262364264
    cout << "ln(1+0.3) = 0.262364264 | COMPUTED: " << y << endl;
    
    y = poly(0.7, params); // 0.53062825106
    cout << "ln(1+0.7) = 0.53062825106 | COMPUTED: " << y;
    cout << endl;

    //////////////////////////////////////////////////////

    vector<double> paramsAT = arcTan(1024*howManyTerms);
    cout << "TERMS FOR ARCTAN :" << endl;
    for (int k = 0; k < howManyTerms ; k++) {
        cout << "\t" << paramsAT[k] << endl;
    }
    y = poly(1.0, paramsAT); // 0.785398163
    cout << "arctan(1) = 0.785398163 | COMPUTED: " << y << endl;
    y = poly(0.577350269, paramsAT); // 0.523598775
    cout << "arctan(0.577350269) = 0.523598775 | COMPUTED: " << y << endl;
}

void DFTandFFTtermsCompare(){
    int n = 8; // always a power of 2
    vector<double> a;
    for (int k = 0 ; k < n; k++) {
        a.push_back((double)(k+1));
    }

    vector<complex<double>> aFFT, aDFT;

    aFFT = FFT(a);
    aDFT = DFT(a);

    for (int k = 0 ; k < n ; k++) {
        cout << "DFT[" << k << "] = " << aDFT[k] << " | FFT[" << k << "] = " << aFFT[k] << endl;
    }
}

void DFTandFFT(){

    int howManyTerms;
    vector<double> paramsDFT, paramsFFT;
    vector<complex<double>> aFFT, aDFT;

    for (int n = 8 ; n <= 15 ; n++) {
        
        // create polynomial parameters
        howManyTerms = pow(2,n);
        paramsDFT = naturalLogarithm(howManyTerms);
        paramsFFT = naturalLogarithm(howManyTerms);

        // NOW COMPUTE

        // FFT
        auto startFFT = high_resolution_clock::now();
        aFFT = FFT(paramsFFT);
        auto stopFFT = high_resolution_clock::now();

        // DFT
        auto startDFT = high_resolution_clock::now();
        aDFT = DFT(paramsDFT);        
        auto stopDFT = high_resolution_clock::now();


        auto durationDFT = duration_cast<microseconds>(stopDFT - startDFT);
        auto durationFFT = duration_cast<microseconds>(stopFFT - startFFT);

        cout << "Transformation for polynomial of " << howManyTerms << " terms:" << endl;
        cout << "\tComputing DFT took " << durationDFT.count() << " microseconds" << endl;
        cout << "\tComputing FFT took " << durationFFT.count() << " microseconds" << endl;
    }

}