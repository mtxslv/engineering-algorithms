#include <iostream>
#include <vector>
#include "lup.h"
#include "processFile.h"

using namespace std;


int main(int argc, char* argv[]) {
    if (argc != 3) {
        cerr << "Usage: " << "lupinverse <matrix_file>" << endl;
        return 1;
    }

    vector<vector<float>> T = readFile(argv[1]); // Tableau matrix

    if (!checkMatrixSquare(T)) {
        cout << "Matrix must be square" << endl;
        return 1;
    }

    // simplexTableau(T,n,m);
    // cout << "Processed tableau containing "<< n << " variable(s) and " << m << " constraint(s)." << endl;
    // string ans = processOutput(T,n,m);
    // cout << ans;
}