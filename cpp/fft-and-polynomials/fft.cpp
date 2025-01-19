#include <complex>
#include <vector>

using namespace std;

vector<complex<double>> DFT(vector<double> a){
    int n = a.size();
    vector<complex<double>> y = vector<complex<double>>(n,complex<double>(0,0));
    
    complex<double> unity_root_k;
    for (int k = 0; k < n ; k++) {
        unity_root_k = polar(1.0,(double)n/(double)k);
        for (int j = 0; j < n; j++) {
            // y[k] += a[j]
        }
    }
    return y;
}

vector<complex<double>> FFT(vector<double> a){
    int n = a.size();
    vector<complex<double>> y = vector<complex<double>>(n,complex<double>(0,0));
    return y;
}


//https://www.reddit.com/r/DSP/comments/13ji73d/question_about_c_implementation_of_dft/
// https://stackoverflow.com/questions/51679516/discrete-fourier-transform-c