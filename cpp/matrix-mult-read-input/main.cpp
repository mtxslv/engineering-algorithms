#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <sstream>  // for stringstream
using namespace std;

void readMatrixFromFile(const char* filePath, vector<vector<float>>& m1, vector<vector<float>>& m2) {
    ifstream file(filePath);
    string line;
    bool readingMatrix1 = true;
    
    if (file.is_open()) {
        // Read the first matrix
        while (getline(file, line)) {
            if (line.empty()) continue;  // Skip empty lines
            stringstream ss(line);
            vector<float> row;
            float value;
            while (ss >> value) {
                row.push_back(value);
                if (ss.peek() == ',') ss.ignore();  // Ignore the comma
            }
            if (readingMatrix1) {
                m1.push_back(row);
            } else {
                m2.push_back(row);
            }
            
            // Check if it's time to switch to the second matrix (after 8 lines)
            if (m1.size() == 8 && readingMatrix1) {
                readingMatrix1 = false;
            }
        }
        file.close();
    } else {
        cout << "Unable to open file." << endl;
    }
}

void mulMat(vector<vector<float>>& m1, vector<vector<float>>& m2, 
            vector<vector<float>>& res) {
    int r1 = m1.size();
    int c1 = m1[0].size();
    int r2 = m2.size();
    int c2 = m2[0].size();

    if (c1 != r2) {
        cout << "Invalid Input" << endl;
        exit(EXIT_FAILURE);
    }

    // Resize result matrix to fit the result dimensions
    res.resize(r1, vector<float>(c2, 0)); 
  
    for (int i = 0; i < r1; i++) {
        for (int j = 0; j < c2; j++) {
            for (int k = 0; k < c1; k++) {
                res[i][j] += m1[i][k] * m2[k][j];
            }
        }
    }
}

// Driver code
int main(int argc, const char * argv[]) {

    if (argc < 2) {
        cout << "No file path." << endl;
        return 1;
    }

    const char* filePath = argv[1];
    vector<vector<float>> m1, m2;

    readMatrixFromFile(filePath, m1, m2);

    // Output matrices (for verification)
    cout << "Matrix 1:" << endl;
    for (const auto& row : m1) {
        for (const auto& value : row) {
            cout << value << " ";
        }
        cout << endl;
    }

    cout << "Matrix 2:" << endl;
    for (const auto& row : m2) {
        for (const auto& value : row) {
            cout << value << " ";
        }
        cout << endl;
    }
    
    vector<vector<float>> res;
    
    mulMat(m1,m2,res);
    cout << "Multiplication of given two matrices is:\n";
    for (const auto& row : res) {
        for (int val : row) {
            cout << val << "\t";
        }
        cout << endl;
    }    

    return 0;
}
