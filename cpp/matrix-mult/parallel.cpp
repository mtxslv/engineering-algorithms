#include <iostream>
#include <vector>
#include <omp.h>

using namespace std;

void mulMat(vector<vector<int>>& m1, vector<vector<int>>& m2, 
            vector<vector<int>>& res) {
    int r1 = m1.size();
    int c1 = m1[0].size();
    int r2 = m2.size();
    int c2 = m2[0].size();

    if (c1 != r2) {
        cout << "Invalid Input" << endl;
        exit(EXIT_FAILURE);
    }

    // Resize result matrix to fit the result dimensions
    res.resize(r1, vector<int>(c2, 0)); 
    
    #pragma omp for
    for (int i = 0; i < r1; i++) {
        for (int j = 0; j < c2; j++) {
            for (int k = 0; k < c1; k++) {
                res[i][j] += m1[i][k] * m2[k][j];
            }
        }
    }
}

// Driver code
int main() {
    vector<vector<int>> m1 = { 
        {1, 2, 0, 0}, 
        {3, 4, 0, 0}, 
        {0, 0, 1, 2}, 
        {0, 0, 3, 4} 
    };
    vector<vector<int>> m2 = { 
        {5, 6, 0, 0}, 
        {7, 8, 0, 0}, 
        {0, 0, 5, 6}, 
        {0, 0, 7, 8} 
    };
    vector<vector<int>> res;

    mulMat(m1, m2, res);

    cout << "Multiplication of given two matrices is:\n";
    for (const auto& row : res) {
        for (int val : row) {
            cout << val << "\t";
        }
        cout << endl;
    }

    return 0;
}
