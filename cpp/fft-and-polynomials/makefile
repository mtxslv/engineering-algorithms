CC=g++
CFLAGS=-I.
DEPS = fft.h taylorSeries.h

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

make: main.o fft.o taylorSeries.o
	$(CC) -o polynomialfft main.o fft.o taylorSeries.o
	rm *.o

simpleclean:
	rm *.o