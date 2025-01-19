#include <complex>
#include <vector>

#include <iostream>

using namespace std;

/* 
    Return the n roots of unity in cartesian form.
    For more information, read:
        https://en.cppreference.com/w/cpp/numeric/complex/polar
*/
vector<complex<double>> rootsOfUnity(int n) {
    vector<complex<double>> roots(n, complex<double>(0, 0));
    double phase;
    for (int k = 0; k < n; k++) {
        phase = -2 * M_PI * k / n; // Negative sign for forward DFT
        roots[k] = polar(1.0, phase); // Magnitude 1, phase angle
    }
    return roots;
}  

vector<complex<double>> DFT(vector<double> a){
    int n = a.size();
    vector<complex<double>> roots = rootsOfUnity(n);

    vector<complex<double>> y(n, complex<double>(0, 0));
    for (int k = 0; k < n; k++) {
        for (int j = 0; j < n; j++) {
            y[k] += a[j] * roots[(k * j) % n]; // Ensure index wraps around
        }
    }
    return y;
}

vector<complex<double>> FFT(vector<double> a){
    int n = a.size();
    vector<complex<double>> y;
    
    if (n == 1){
        y.resize(1);
        y[0] = complex<double>(a[0],0); // DFT of 1 element is the element itself
        return y;
    }
    
    y.resize(n);

    complex<double> omega_n = polar(1.0,-2*M_PI/(double)n);
    complex<double> omega = complex<double>(1,0);

    // split vector a into even and odd
    vector<double> aEven, aOdd;
    for (int k; k < n; k++){
        if (k%2 == 0) { // k even 
            aEven.push_back(a[k]);
        } else {
            aOdd.push_back(a[k]);
        }
    }

    vector<complex<double>> yEven = FFT(aEven);
    vector<complex<double>> yOdd = FFT(aOdd);

    for(int k = 0; k <= n/2-1; k++){
        y[k] = yEven[k] + omega*yOdd[k];
        y[k+n/2] = yEven[k] - omega*yOdd[k];
        omega = omega*omega_n;
    }

    return y;
}


//https://www.reddit.com/r/DSP/comments/13ji73d/question_about_c_implementation_of_dft/
// https://stackoverflow.com/questions/51679516/discrete-fourier-transform-c