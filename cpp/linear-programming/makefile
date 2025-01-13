CC=g++
CFLAGS=-I.
DEPS = simplex.h processFile.h userInput.h

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

make: main.o simplex.o processFile.o userInput.o
	$(CC) -o simplexsolver main.o simplex.o processFile.o userInput.o
	rm *.o

simpleclean:
	rm *.o