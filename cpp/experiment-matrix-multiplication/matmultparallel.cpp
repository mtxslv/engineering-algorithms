#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <chrono>
#include <omp.h>

using namespace std;
using namespace chrono;

vector<vector<int>> read_matrix_from_file(const string& filename) {
    vector<vector<int>> matrix;

    ifstream file(filename);
    if (!file.is_open()) {
        cerr << "Error opening file: " << filename << endl;
        return matrix;
    }

    string line;
    while (getline(file, line)) {
        vector<int> row;
        istringstream iss(line);
        int value;

        while (iss >> value) {
            row.push_back(value);
        }

        matrix.push_back(row);
    }

    file.close();
    return matrix;
}

void mulMat(vector<vector<int>>& m1, vector<vector<int>>& m2, 
            vector<vector<int>>& res) {
    int r1 = m1.size();
    int c1 = m1[0].size();
    int r2 = m2.size();
    int c2 = m2[0].size();

    if (c1 != r2) {
        cout << "Invalid Input" << endl;
        exit(EXIT_FAILURE);
    }

    // Resize result matrix to fit the result dimensions
    res.resize(r1, vector<int>(c2, 0)); 
  
    #pragma omp parallel for
    for (int i = 0; i < r1; i++) {
        for (int j = 0; j < c2; j++) {
            for (int k = 0; k < c1; k++) {
                res[i][j] += m1[i][k] * m2[k][j];
            }
        }
    }
}

// Driver code
int main(int argc, char* argv[]) {
    if (argc != 5) {
        cerr << "Usage: " << argv[0] << " <first_matrix_file> <second_matrix_file> <output_file> <log_file>" << endl;
        return 1;
    }

    string firstMatrixFilePath = argv[1];
    string secondMatrixFilePath = argv[2];
    string outputFilePath = argv[3];
    string logFilePath = argv[4];
    
    vector<vector<int>> m1 = read_matrix_from_file(firstMatrixFilePath);
    vector<vector<int>> m2 = read_matrix_from_file(secondMatrixFilePath);

    vector<vector<int>> res;

    auto start = high_resolution_clock::now();
    mulMat(m1, m2, res);
    auto stop = high_resolution_clock::now();

    auto duration = duration_cast<microseconds>(stop - start);

    // Open the output file
    ofstream outfile(outputFilePath);
    if (!outfile.is_open()) {
        cerr << "Error opening output file: " << outputFilePath << endl;
        return 1;
    }

    // Write the result matrix to the file
    for (const auto& row : res) {
        for (int val : row) {
            outfile << val << "\t";
        }
        outfile << endl;
    }

    outfile.close();

    cout << "Result saved to: " << outputFilePath << endl;
    cout << "Execution time: " << duration.count() << " microseconds" << endl;

    // Open the output file
    ofstream logfile(logFilePath);
    if (!logfile.is_open()) {
        cerr << "Error opening log file: " << outputFilePath << endl;
        return 1;
    }    
    
    logfile << "Execution time: " << duration.count() << " microseconds" << endl;

    logfile.close();
    return 0;

}
