#include <math.h>       /* pow */
#include <vector>

using namespace std;

double poly(double x, vector<double> params){
    double y = 0.0;
    int n = params.size();
    for (int k = 0; k < n; k++) {
        y += params[k]*pow(x,k);
    }
    return y;
}