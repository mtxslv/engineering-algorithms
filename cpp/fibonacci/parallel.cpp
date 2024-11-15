// COMPILE IT RUNNING
// $ g++ -o parallel -fopenmp parallel.cpp

#include<iostream>
#include <omp.h>

using namespace std;

int p_fib(int n) {
    if (n <= 1) {
        return n;
    } else {
        int x, y;
        #pragma omp task shared(x)
        x = p_fib(n-1); // SPAWN HERE

        #pragma omp task shared(y)
        y = p_fib(n-2);

        #pragma omp taskwait // SYNC HERE

        return x + y;
    }
}

int main(){
    int x = 4, fib_4;

    #pragma omp parallel
    {
        #pragma omp single
        {
            fib_4 = p_fib(x);
        }
    }

    cout << "fib(4) = " << fib_4 << endl;

    return 0;
}
