#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>

using namespace std;

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

int main(int argc, char* argv[]) {
    if (argc != 2) {
        cerr << "Usage: " << argv[0] << " <filename>" << endl;
        return 1;
    }

    string filename = argv[1];
    vector<vector<int>> matrix = read_matrix_from_file(filename);

    // Print the matrix (optional)
    for (const auto& row : matrix) {
        for (int value : row) {
            cout << value << " ";
        }
        cout << endl;
    }

    return 0;
}