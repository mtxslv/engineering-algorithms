#include <iomanip>
#include <iostream>
#include <cstring>
#include <vector>
#include "fft.h"
#include "polynomial.h"
#include "taylorSeries.h"

using namespace std;

int main(int argc, char* argv[]) {
    vector<double> params;
    double x = 0.5;
    double y ;
    cout << "NATURAL LOGARITHM OF " << 1 + x << endl;
    
    int terms = 100+1 ;
    for(int k = 10; k <= terms; k++){
        params = naturalLogarithm(k);
        y = poly(x,params);
        cout << "\t" << k <<" params: " ;
        cout << setprecision(25) << y << endl;
    }   
}