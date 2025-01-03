#include <iostream>
#include <string.h>
#include <vector>
#include "lup.h"
#include "matmul.h"
#include "processFile.h"

using namespace std;

int main(int argc, char* argv[]) {
    if (argc != 2 && argc != 4) {
        cerr << "Usage: " << "./lupinverse <matrix_file> -c <inverse_file>" << endl;
        cerr << "The -c flag and its argument is optional for checking if the product is the identity matrix." << endl;
        return 1;
    }

    if (argc == 2){
        vector<vector<float>> A = readFile(argv[1]);

        if (!checkMatrixSquare(A)) {
            cout << "Matrix must be square" << endl;
            return 1;
        }

        vector<vector<float>> Inverse = LupInverse(A);

        printMatrix(Inverse);
    } else {
        if (strcmp(argv[2], "-c")){
            cerr << "Usage: " << "./lupinverse <matrix_file> -c <inverse_file>" << endl;
            cerr << "The -c flag and its argument is optional for checking if the product is the identity matrix." << endl;
            return 1;            
        }

        vector<vector<float>> A = readFile(argv[1]);
        vector<vector<float>> O = readFile(argv[3]);

        if (!checkMatrixSquare(A) || !checkMatrixSquare(O)) {
            cerr << "Both matrices must be square" << endl;
            return 1;
        }

        if (A.size() != O.size()) {
            cerr << "Matrices must have the same dimensions" << endl;
            return 1;
        }

        vector<vector<float>> identity;
        mulMat(A,O,identity);
        printMatrix(identity);
    }

}