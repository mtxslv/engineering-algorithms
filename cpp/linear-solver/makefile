CC=g++
CFLAGS=-I.
DEPS = file1.h

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

make: main.o file1.o
	$(CC) -o main main.o file1.o

simpleclean:
	rm *.o