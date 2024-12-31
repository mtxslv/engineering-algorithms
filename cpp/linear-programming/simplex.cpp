#include <iostream>
#include <vector>
#include "simplex.h"

using namespace std;

bool checkIfNegative(vector<vector<float>>& T, int n){
    for (int i = 1; i < n+1 ; i++) { // n variables. T[0][0] is Z
        if (T[0][i] < 0) {
            return true;
        }
    }
    return false;
}


void simplexTableau(
    vector<vector<float>>& T,
    int n,
    int m
) {
    vector<float> q;


}