// COMPILE IT RUNNING
// $ g++ -o parallel -fopenmp parallel.cpp

#include<iostream>
#include <omp.h>
#include <string>

using namespace std;

int p_fib(int n) {
    if (n <= 1) {
        return n;
    } else {
        int x, y;
        #pragma omp task shared(x)
        x = p_fib(n-1); // SPAWN HERE

        y = p_fib(n-2); 

        #pragma omp taskwait // SYNC HERE

        return x + y;
    }
}

int processInput(int argc, char * argv[]){
    return std::stoi(argv[1]);
}

int main(int argc, char * argv[]){

    if (argc == 1) {
        cout << "No argument provided" <<endl;
        return 0;
    } else {

        int x = processInput(argc, argv);
        int fib;

        #pragma omp parallel
        {
            fib = p_fib(x);
        }

        cout << "FIBONACCI(" << x << ") = "<< fib << endl;

        return 0;
    }
}
