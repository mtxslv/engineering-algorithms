CC=g++
CFLAGS=-I.
DEPS = fft.h polynomial.h taylorSeries.h experiments.h

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

make: main.o fft.o polynomial.o taylorSeries.o experiments.o
	$(CC) -o polynomialfft main.o fft.o polynomial.o taylorSeries.o experiments.o
	rm *.o

simpleclean:
	rm *.o