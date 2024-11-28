#include <omp.h>
#include <iostream>

using namespace std;

int main(int argc, char* argv[]){
    if (argc<1) {
        cerr << "Usage: " << argv[0] << " $OMP_NUM_THREADS" <<endl;
        return 1;
    }

    cout << "SUGGESTED NUM THREADS: "<< argv[1] << endl;

	int nthreads, tid;

	// Begin of parallel region
	#pragma omp parallel private(nthreads, tid)
	{
		// Getting thread number
		tid = omp_get_thread_num();
		printf("Welcome to GFG from thread = %d\n",
			tid);

		if (tid == 0) {

			// Only master thread does this
			nthreads = omp_get_num_threads();
			printf("Number of threads = %d\n",
				nthreads);
		}
	}

}

// https://www.geeksforgeeks.org/openmp-hello-world-program/