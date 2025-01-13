#include <iostream>
#include <cstring>
#include "simplex.h"
#include "processFile.h"
#include "userInput.h"


using namespace std;

int main(int argc, char* argv[]) {

    if (argc != 1 && argc != 5) {
        cerr << "You can call the interactive window with:" << endl;
        cerr << "\t ./simplexsolver" << endl;
        cerr << "Optionally, you can directly provide a tableau file using: " << endl;
        cerr << "\t./simplexsolver -f <tableau_file> <number_of_variables> <number_of_constraints>" << endl;
        return 1;
    }

    // first define variables
    vector<vector<float>> T ;
    int n; 
    int m; 

    if (argc == 1) {
        cout << "YOU NEED TO CALL USER INPUT HERE" << endl;
        return 1;
    } else {
        if (strcmp(argv[1], "-f")){
            cerr << "You can call the interactive window with:" << endl;
            cerr << "\t ./simplexsolver" << endl;
            cerr << "Optionally, you can directly provide a tableau file using: " << endl;
            cerr << "\t./simplexsolver -f <tableau_file> <number_of_variables> <number_of_constraints>" << endl;
            return 1;            
        }
        n = atoi(argv[3]); // number of variables
        m = atoi(argv[4]); // number of constraints

        T = readFile(argv[2]); // Tableau matrix
    }

    if (!checkMatrixFormat(T,n,m)) {
        cout << "Matrix format does not conform to quantity of variables and restrictions" << endl;
        return 1;
    }

    if (isMinimizationProblem(T)){
        // METHOD TO GET FIRST VALID SOLUTION
        firstFeasibleSolutionMinimization(T,n,m);
        // TELL IT IS MINIMIZATION PROBLEM
        cout << "Minimization problem" << endl;
    } else {
        cout << "Maximization problem" << endl;
    }
    simplexTableau(T,n,m);
    cout << "Processed tableau containing "<< n << " variable(s) and " << m << " constraint(s)." << endl;
    string ans = processOutput(T,n,m);
    cout << ans;
}

/*
   1    -40    -30    0    0    0
   0      1      1    1    0   12
   0      2      1    0    1   16

    vector<vector<float>> T {
        {1.0f, -40.0f, -30.0f, 0.0f, 0.0f,  0.0f},
        {0.0f,   1.0f,   1.0f, 1.0f, 0.0f, 12.0f},
        {0.0f,   2.0f,   1.0f, 0.0f, 1.0f, 16.0f}
    }; 
    simplexTableau(T,2,2); // 2 variables and 2 constraints
    printMatrix(T);
    cout << endl << endl << "FINAL MATRIX::" << endl;
    printMatrix(T);
*/

/*
    vector<vector<float>> T {
        {1.0f, -4.0f, -3.0f, 0.0f, 0.0f, 0.0f, 0.0f, 0.0f},
        {0.0f,  2.0f,  3.0f, 1.0f, 0.0f, 0.0f, 0.0f, 6.0f},
        {0.0f, -3.0f,  2.0f, 0.0f, 1.0f, 0.0f, 0.0f, 3.0f},
        {0.0f,  0.0f,  2.0f, 0.0f, 0.0f, 1.0f, 0.0f, 5.0f},
        {0.0f,  2.0f,  1.0f, 0.0f, 0.0f, 0.0f, 1.0f, 4.0f}
    }; 
    simplexTableau(T,2,4); // 2 variables and 4 constraints
    printMatrix(T);
    cout << endl << endl << "FINAL MATRIX::" << endl;
    printMatrix(T);
*/