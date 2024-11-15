#include <iostream>
#include <vector>
using namespace std;

void simpleVecMatMul(vector<vector<float>>& mat, vector<float>& vec, vector<float>& res) {
    
    // Get dimensions
    int rowsMatrix = mat.size();
    int colsMatrix = mat[0].size();

    int rowsVector = vec.size();
    
    // Make sure multiplication is valid
    if (colsMatrix != rowsVector) {
        cout << "Invalid Input" << endl;
        exit(EXIT_FAILURE);
    } else {
        cout << "MATRIX SHAPE: ("<< rowsMatrix <<","<<colsMatrix << ")" << endl;
        cout << "VECTOR SHAPE: ("<< rowsVector <<"," << ")" << endl;
    }

    // Make sure res has right dimension to store result
    res.resize(rowsMatrix, 1);
    for (int k = 0; k < rowsMatrix; k++){
        res[k] = 0;
    }
    cout << "OUTPUT VECTOR SHAPE: (" << res.size() << ",)"<< endl;

    // Multiply using loop
    for (int i=0; i < rowsMatrix; i++){ // For each matrix row
        // cout << "Row "<<i<<endl;
        for (int k=0; k < colsMatrix; k++){
            res[i] += mat[i][k]*vec[k];
            // cout << mat[i][k] << "*" << vec[k]<< endl;
        }
    }
}

// Driver code
int main() {
    vector<vector<float>> m1 = { 
        {1, 2}, 
        {3, 4}
    };
    vector<float> v = { 
        {5}, 
        {7} 
    };
    vector<float> res;

    simpleVecMatMul(m1, v, res);

    int el = res.size();
    cout << "Multiplication of given two matrices is:\n";
    for (int i = 0; i < el; i++) {
        cout << res[i] << "\n";
    }
    return 0;
}
