#include <math.h>       /* pow */
#include "taylorSeries.h"

using namespace std;

/*
    Actually, naturalLogarithm() implements
    the parameters for the series of
            ln(1+x)
*/
vector<double> naturalLogarithm(int n){
    vector<double> params = vector<double>(n+1, 0.0); 
    for (int k=1; k<=n ; k++){
        params[k] = pow(-1,k-1)/k;
    }
    return params;
}


// vector<double> arcTan(int n) {

// }