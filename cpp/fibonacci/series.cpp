#include<iostream>

using namespace std;

int fib(int n) {
    if (n <= 1) {
        return n;
    } else {
        int x = fib(n-1);
        int y = fib(n-2);
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
        int ans;

        ans = fib(x);

        cout<<"FIBONACCI("<<x<< ") = "<<ans<< endl;

        return 0;
    }
}