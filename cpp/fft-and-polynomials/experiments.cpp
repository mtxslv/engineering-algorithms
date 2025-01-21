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

void convolution1DTest(){
    vector<complex<double>> A = {
        complex<double>(1.0, 0.0),
        complex<double>(2.0, 0.0),
        complex<double>(3.0, 0.0),
    };
    vector<complex<double>> B = {
        complex<double>(0.0, 0.0),
        complex<double>(1.0, 0.0),
        complex<double>(0.5, 0.0),
    };    
    vector<complex<double>> C = conv1D(A,B);
    for (int i = 0; i < C.size(); i++){
        cout << C[i] << endl;
    }
}

void regularPolynomialProduct(){
    // https://runestone.academy/ns/books/published/int-algebra/PolynomialFunctions.html

    vector<double> A = {2.0, 1.0}; // x + 2
    vector<double> B = {4.0, 0.0, -3.0, 5.0}; // 4 + 0x -3x² + 5x³ 
    vector<double> C = polynomialProduct(A,B); // 8 + 4x -6x² + 7x³ + 5x⁴

    cout << "C(x) = " << C[0] ;
    for(int i = 1; i < C.size(); i++){
        if (C[i] >= 0) {
            cout << " + " ;
        } else {
            cout << " ";
        }
        cout << C[i] << "*x^" << i;
    }
    cout << endl;
}

void FFTandIDFT(){
    vector<double> A = {1.0, 2.0, 3.0, 4.0};
    // vector<double> A = {7.0, 2.0, 9.0, 5.0};
    vector<complex<double>> fftA = FFT(A); // 

    cout << "FFT COEFFICIENTS: " << endl;
    for (int k = 0; k < fftA.size(); k++){
        cout << fftA[k] << "\t";
    }
    cout << endl;

    vector<complex<double>> restoredA = IDFT(fftA);


    cout << "REGULAR COEFFICIENTS: " << endl;
    for (int k = 0; k < restoredA.size(); k++){
        cout << restoredA[k] << "\t";
    }
    cout << endl;

    cout << "COEFFICIENTS MUST BE: 1, 2, 3, 4" << endl;


}

void comparisonFFTandPolyProd(){
    vector<double> A = {2.0, 1.0}; // x + 2
    vector<double> B = {4.0, 0.0, -3.0, 5.0}; // 4 + 0x -3x² + 5x³ 

    // Print both polynomials
    cout << endl << "Polynomials to multiply:" << endl;
    printPoly('A', A);
    printPoly('B', B);

    // Regular poly prod
    vector<double> Cconv1D = conv1D(A,B);
    vector<double> C = polynomialProduct(A,B); // 8 + 4x -6x² + 7x³ + 5x⁴


    cout << endl << "Polynomial from regular term-by-term product: " << endl;
    printPoly('C', C);
    cout << "Polynomial from 1D Convolution: " << endl;
    printPoly('C', Cconv1D);

    // Pad A to have 4 points
    A.push_back(0);
    A.push_back(0);

    // Now pad both to 2n
    for (int k = 0; k < 4; k++){
        A.push_back(0);
        B.push_back(0);
    }

    cout << endl << "We pad both parameters vector to have the same power-2 size..." << endl;
    cout << "A parameters: [\t" ;
    for (int i = 0; i < A.size() ; i++){
        cout << A[i] << "\t" ;
    }
    cout << "]" << endl;

    cout << "B parameters: [\t" ;
    for (int i = 0; i < B.size() ; i++){
        cout << B[i] << "\t" ;
    }
    cout << "]" << endl;    

    // FFT each polynomial
    vector<complex<double>> aFFT = FFT(A);
    vector<complex<double>> bFFT = FFT(B);

    cout << endl << "Now we take the FFT of each polynomial's parameter vector..." << endl;
    cout << "fft{A} = [\t";
    for (int i = 0; i < aFFT.size() ; i++){
        cout << aFFT[i].real() ;
        if(aFFT[i].imag() < 0) {
           cout <<  aFFT[i].imag();
        } else {
           cout << "+ " << aFFT[i].imag();
        }    
        cout << "i\t";
    }
    cout << "]" << endl;
    cout << "fft{B} = [\t";
    for (int i = 0; i < bFFT.size() ; i++){
        cout << bFFT[i].real() ;
        if(bFFT[i].imag() < 0) {
           cout <<  bFFT[i].imag();
        } else {
           cout << "+ " << bFFT[i].imag();
        }    
        cout << "i\t";
    }
    cout << "]" << endl;    

    // pointwise prod in frequency domain
    vector<complex<double>> cFFT = pointwiseProd(aFFT, bFFT);
    cout << "The pointwise product creates a new parameter vector... let's call it..." << endl;
    cout << "fft{C} = [\t";
    for (int i = 0; i < cFFT.size() ; i++){
        cout << cFFT[i].real() ;
        if(cFFT[i].imag() < 0) {
           cout <<  cFFT[i].imag();
        } else {
           cout << "+ " << cFFT[i].imag();
        }    
        cout << "i\t";
    }
    cout << "]" << endl;   

    // // IDFT the result
    vector<complex<double>> recC = IDFT(cFFT);
    cout << endl << "We IDFT the result, creating a 'restored' C: " << endl;
    cout << "IDFT{fft{C}} = [\t";
    for (int i = 0; i < recC.size() ; i++){
        if (recC[i].real() < 1e-15 && recC[i].real() > -1e-15){
            cout << 0 ;
        } else {
            cout << recC[i].real() ;
        }
        if(recC[i].imag() < 1e-15 && recC[i].imag() > -1e-15) {
           cout <<  "+0i\t"; 
        }    
    }
    cout << "]" << endl;   

    cout<< endl << "Notice these are the parameters of the product C!"<< endl;

}