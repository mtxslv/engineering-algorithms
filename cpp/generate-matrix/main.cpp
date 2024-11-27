// compile with
// $ g++ -o main main.cpp
// generate matrices using:
// $ ./main 5 1223334444 ./matrix_a.txt
// $ ./main 5 4444333221 ./matrix_b.txt

#include <iostream>
#include <string>
#include <cstdlib>
#include <fstream>
#include <stdexcept>

using namespace std;

unsigned int castToUnsignedInt(const std::string& s)
{
    unsigned long lresult = stoul(s, 0, 10);
    unsigned int result = lresult;
    if (result != lresult) throw std::out_of_range("OutOfRangeError");
    return result;
}

int main(int argc, char * argv[]) {

    if (argc < 4) {
        cout << "No argument provided" << endl;
        return 0;
    } else {

        int n = std::stoi(argv[1]);
        unsigned int seed = castToUnsignedInt(argv[2]);
        string path = argv[3];

        cout << "Generating " << n << " terms (using ";
        cout << "uint "<< seed<< " as seed) and saving to: \n\t" << path << endl;

        // // Seed the random number generator with the current time
        // srand(seed);

        // // Open the file for writing
        // ofstream outfile(path);

        // if (!outfile.is_open()) {
        //     cerr << "Error opening file: " << path << endl;
        //     return 1;
        // }

        // // Generate n^2 random integers and write them to the file
        // for (int i = 0; i < n * n; ++i) {
        //     int random_number = rand() % 256;
        //     outfile << random_number << " ";

        //     // Add a newline after every n values
        //     if ((i + 1) % n == 0) {
        //         outfile << endl;
        //     }
        // }

        // outfile.close();

        // cout << "File saved successfully!" << endl;
        // return 0;
    }
}