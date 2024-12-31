CC=g++
CFLAGS=-I.
DEPS = simplex.h

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

make: main.o simplex.o processFile.o
	$(CC) -o main main.o simplex.o processFile.o
	rm *.o

simpleclean:
	rm *.o