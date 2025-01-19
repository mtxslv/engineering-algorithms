#include <iomanip>
#include <iostream>
#include <cstring>
#include <vector>
#include "fft.h"
#include "polynomial.h"
#include "taylorSeries.h"

using namespace std;

int main(int argc, char* argv[]) {
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