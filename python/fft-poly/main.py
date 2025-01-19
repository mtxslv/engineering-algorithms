import numpy as np

A = np.array([1.0, 2.0, 3.0, 4.0, 5.0])
n = len(A)
N = 1 << (n - 1).bit_length()  # Next power of two
# A_padded = np.pad(A, (0, N - len(A)))
A_fft = np.fft.fft(A)
print(A_fft)