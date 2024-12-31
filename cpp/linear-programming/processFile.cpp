#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>
#include <string>
#include "processFile.h"

using namespace std;

vector<vector<float>> readFile(const char * filePath){
    vector<vector<float>> T ;

    ifstream file(filePath);
    if (!file.is_open()) {
        cerr << "Error opening file: " << filePath << endl;
        return T;
    }

    string line;
    while (getline(file, line)) {
        vector<float> row;
        istringstream iss(line);
        float value;
        while(iss >> value) {
            row.push_back(value);
        }
        T.push_back(row);
    }

    return T;
}

bool checkMatrixFormat(vector<vector<float>>& T, int n, int m) {
    int howManyRows = T.size();

    if (howManyRows != m+1) { // m restrictions and function to maximize
        return false;
    }

    for( int i = 0 ; i <  howManyRows; i++) {
        if(T[i].size() != n+m+2){ // n variables, m restrictions, z an b cols
            return false;
        }
    }
    
    return true;
}