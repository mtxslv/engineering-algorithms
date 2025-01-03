#include <iostream>
#include <vector>
#include "lup.h"
#include "processFile.h"

using namespace std;


int main(int argc, char* argv[]) {
    if (argc != 2 && argc != 4) {
        cerr << "Usage: " << "./lupinverse <matrix_file> -c <inverse_file>" << endl;
        cerr << "The -c flag and its argument is optional for checking if the product is the identity matrix." << endl;
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