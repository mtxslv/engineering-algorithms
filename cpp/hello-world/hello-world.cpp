#include <iostream>
using namespace std;

void printMessage() {

}

int main(int argc, const char * argv[]) {
    cout << argc << " args were passed. They are:\n";
    for(int i=0; i < argc; i++){
        std::cout << argv[i] << "\n";
    }
    cout << endl;
}