#include<iostream>

using namespace std;

int fib(int n) {
    if (n <= 1) {
        return 1;
    } else {
        int x = fib(n-1);
        int y = fib(n-2);
        return x + y;
    }
}

int main(){
    int x = 4;
    int fib_4 = fib(4);
    cout << "fib(4) = " << fib_4 << endl;
}