#include <fstream>
#include <iostream>
#include <string>

using namespace std;

void processLine(string line) {
    int charNumber = line.size();
    cout << line << " ~~~> " << charNumber << endl;
}

void readFile(const char * filePath) {
    ifstream file(filePath);
    string line;

    if (file.is_open()){
        while (getline(file, line)) {
            processLine(line);
        }
        file.close();
    } 

}

int main(int argc, const char * argv[]){
    if (argc < 2) {
        cout << "No file path.";
    } else {
        const char* filePath = argv[1];
        readFile(filePath);
    }
    return 0;
}