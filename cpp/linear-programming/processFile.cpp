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

    return T;


}