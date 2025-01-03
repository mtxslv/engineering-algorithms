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

bool checkMatrixSquare(vector<vector<float>>& T) {
    int howManyRows = T.size();

    for( int i = 0 ; i <  howManyRows; i++) {
        if(T[i].size() != howManyRows){ // n variables, m restrictions, z an b cols
            return false;
        }
    }
    
    return true;
}