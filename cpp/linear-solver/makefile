CC=g++
CFLAGS=-I.
DEPS = file1.h processFile.h lup.h matmul.h

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

make: main.o file1.o processFile.o lup.o matmul.o
	$(CC) -o lupinverse -fopenmp main.o file1.o processFile.o lup.o matmul.o
	rm *.o

simpleclean:
	rm *.o