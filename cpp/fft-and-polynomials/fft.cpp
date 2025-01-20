#include <bit>
#include <cmath>
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

// Compute FFT
vector<complex<double>> FFT(vector<double> a) {
    int n = a.size();
    vector<complex<double>> y;

    if (n == 1) {
        y.resize(1);
        y[0] = complex<double>(a[0], 0); // DFT of 1 element is the element itself
        return y;
    }

    y.resize(n);

    complex<double> omega_n = polar(1.0, -2 * M_PI / (double)n);
    complex<double> omega = complex<double>(1, 0);

    // Split vector a into even and odd indices
    vector<double> aEven, aOdd;
    for (int k = 0; k < n; k++) {
        if (k % 2 == 0) { // k even
            aEven.push_back(a[k]);
        } else {
            aOdd.push_back(a[k]);
        }
    }

    vector<complex<double>> yEven = FFT(aEven);
    vector<complex<double>> yOdd = FFT(aOdd);

    for (int k = 0; k < n / 2; k++) {
        y[k] = yEven[k] + omega * yOdd[k];
        y[k + n / 2] = yEven[k] - omega * yOdd[k];
        omega = omega * omega_n;
    }

    return y;
}

// Find the next power of 2
int nextPowerOf2(int n) {
    if (n <= 0) {
        return 1; // Or handle as an error
    }
    if ((n & (n - 1)) == 0) { // Already a power of 2
        return n;
    }
    n--;
    n |= n >> 1;
    n |= n >> 2;
    n |= n >> 4;
    n |= n >> 8;
    n |= n >> 16;
    n++;
    return n;
}

// Pad the input to the next power of 2
vector<double> padToPowerOf2(vector<double> a) {
    int n = a.size();
    int N = nextPowerOf2(n);
    a.resize(N, 0.0); // Resize and pad with zeros
    return a;
}

// Safe FFT implementation with padding
vector<complex<double>> safeFFT(vector<double> a) {
    vector<double> okA = padToPowerOf2(a);
    return FFT(okA);
}

// Compute I-FFT
vector<complex<double>> IFFT(vector<complex<double>> a) {
    int n = a.size();
    vector<complex<double>> y;

    if (n == 1) {
        y.resize(1);
        y[0] = a[0]; // DFT of 1 element is the element itself
        return y;
    }

    y.resize(n);

    // Conjugate the input
    for (int i = 0; i < n; i++) {
        a[i] = conj(a[i]);
    }

    // Compute FFT of conjugated input
    complex<double> omega_n = polar(1.0, -2 * M_PI / (double)n); // Note the negative sign
    complex<double> omega = complex<double>(1, 0);

    vector<complex<double>> aEven, aOdd;
    for (int k = 0; k < n; k++) {
        if (k % 2 == 0) {
            aEven.push_back(a[k]);
        } else {
            aOdd.push_back(a[k]);
        }
    }

    vector<complex<double>> yEven = IFFT(aEven);
    vector<complex<double>> yOdd = IFFT(aOdd);

    for (int k = 0; k < n / 2; k++) {
        y[k] = yEven[k] + omega * yOdd[k];
        y[k + n / 2] = yEven[k] - omega * yOdd[k];
        omega = omega * omega_n;
    }

    // Conjugate the result and scale it by dividing by n
    for (int i = 0; i < n; i++) {
        y[i] = conj(y[i]) / (double)n;
    }

    return y;
}