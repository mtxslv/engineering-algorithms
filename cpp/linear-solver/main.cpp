#include <iostream>
#include <vector>
#include "lup.h"
#include "processFile.h"

using namespace std;


int main(int argc, char* argv[]) {
    if (argc != 2) {
        cerr << "Usage: " << "lupinverse <matrix_file>" << endl;
        return 1;
    }

    vector<vector<float>> A = readFile(argv[1]); // Tableau matrix

    if (!checkMatrixSquare(A)) {
        cout << "Matrix must be square" << endl;
        return 1;
    }

    vector<vector<float>> Inverse = LupInverse(A);

    printMatrix(Inverse);
}