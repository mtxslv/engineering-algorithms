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

/*
    arcTan implements the params for arctan(x)
*/
vector<double> arcTan(int n) {
    vector<double> params = vector<double>(n+1, 0.0); 
    for (int k=1; k<=n ; k++){
        if (k%2 == 0) {
            params[k] = 0;
        } else {
            params[k] = pow(-1,k-1)/(2*k-1);
        }
    }
    return params;
} // CHECAR DEPOIS SE TÁ CERTO (ACHO QUE NÃO, NUM SEI)