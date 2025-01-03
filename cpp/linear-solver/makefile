CC=g++
CFLAGS=-I.
DEPS = file1.h processFile.h lup.h

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

make: main.o file1.o processFile.o lup.o
	$(CC) -o main main.o file1.o processFile.o lup.o
	rm *.o

simpleclean:
	rm *.o