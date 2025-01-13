#include <cfloat> // To use FLT MAX
#include <cmath> // To use abs
#include <iostream>
#include <iomanip>
#include <string>
#include <sstream>
#include <vector>
#include "simplex.h"
#include "userInput.h"

using namespace std;

/*
    NEED TO IMPROVE THIS CODE IN A WHILE LOOP
    TO MAKE SURE THE PROPER ANSWERS ARE BEING PROVIDED
*/
void processUserInput(
    std::vector<std::vector<float>>& T, 
    int& n,
    int& m
) {
    bool isMax;
    float ans = -1;

    cout << "Welcome to simplex-solver, by mtxslv! I'm glad to help!" << endl;
    cout << "In order to get started, let me ask you..." << endl << endl; 

    // First question: 
    cout << "Which problem to solve?\n1) Maximization \n2) Minimization" << endl;
    cin >> ans;
    if (ans == 1) {
        isMax = true;
    } else {
        isMax = false;
    }
    // Second question:
    cout << "How many variables?\t" ;
    cin >> n;
    // Third question:
    cout << "How many restrictions?\t";
    cin >> m;
    cout << endl;

    // Now we can start building the matrix
    T.resize(m+1,std::vector<float>(n+m+2));
    // Z
    if (isMax) {
        T[0][0] = 1;
    } else {
        T[0][0] = -1;
    }
    // Slack variables
    for (int k = n+1 ; k < n+m+1 ; k++){
        if (isMax) {
            T[k-n][k] = 1;
        } else {
            T[k-n][k] = -1;
        } 
    }

    // Fourth question
    cout << "Now tell me about your objective function. If it would look like this:\n\t\t z = ";
    for (int k = 0; k < n-1; k++){
        if (isMax) {
            cout << "c" << k+1 << "*x" << k+1 << " + " ;
        } else {
            cout << "b" << k+1 << "*y" << k+1 << " + " ;
        }
    }
    if (isMax) {
        cout << "c" << n << "*x" << n << endl;
    } else {
        cout << "b" << n << "*y" << n << endl;
    }

    // OBJECTIVE FUNCTION VALUES
    cout << "which would be the value of..." << endl;
    for (int k = 0; k < n; k++) {
        if (isMax) {
            cout << "\tc"<< k+1 << ": " ;
            cin >> ans;
            T[0][k+1] = -ans;
        } else {
            cout << "\tb"<< k+1 << ": " ;
            cin >> ans;
            T[0][k+1] = ans;
        }        
    }

    // Fifth question
    cout << "Sweet! Now let's jump to the restrictions. Supposing a generic restriction similar to:";
    cout << "\n\t\t";
    for (int k = 0; k < m-1; k++){
        if (isMax) {
            cout << "a" << k+1 << "*x" << k+1 << " + " ;
        } else {
            cout << "a" << k+1 << "*y" << k+1 << " + " ;
        }
    }
    if (isMax) {
        cout << "a" << m << "*x" << m ;
        cout << " <= b" << endl;
    } else {
        cout << "a" << m << "*y" << m ;
        cout << " >= c" << endl;
    }

    cout << "I'll need you to input the numeric values, so tell me..." << endl;
    for (int i = 0 ; i < m ; i++) {
        cout << "\tthe values for the restriction equation #" << i+1 << "..." << endl;
        for (int j = 0 ; j < n ; j++) {
            cout << "a" << j+1 << ": " ;
            cin >> ans;
            T[i+1][j+1] = ans;
        }
        cout << "... and the independent variable: ";
        cin >> ans;
        T[i+1][n+m+1] = ans;
    }
    cout << "Great! Thank you for your patience! We're done here. I'll proceed to work on your problem!" << endl << endl;
}