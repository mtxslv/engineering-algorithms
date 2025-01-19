from pathlib import Path
import re
from matplotlib import pyplot as plt

data_path = Path(__file__).parents[2] / 'cpp' / 'fft-and-polynomials' / 'dft-and-fft-comparison.txt'
assert data_path.exists() and data_path.is_file()

with open(data_path) as file:
    file_contents = file.readlines()

pattern_trans = r'Transformation for polynomial of (\d.+) terms:'
pattern_DFT = r'Computing DFT took (\d.+) microseconds'
pattern_FFT = r'Computing FFT took (\d.+) microseconds'

k = 0

fft_times = []*8
dft_times = []*8
n = []*8

for it, line in enumerate(file_contents):
    rem = it%3
    if rem == 0: # Transformation
        n_ = int(re.findall(pattern_trans,line)[0])
        n.append(n_)
    elif rem == 1: # DFT
        dft_us = int(re.findall(pattern_DFT,line)[0])
        dft_times.append(dft_us/1e6)
    elif rem == 2: # FFT
        fft_us = int(re.findall(pattern_FFT,line)[0])
        fft_times.append(fft_us/1e6)
        k += 1

fig, (ax1,ax2) = plt.subplots(2,1, figsize=(10, 8), sharex=True)

ax1.plot(n, fft_times, color = 'green', label = "FFT")
ax1.plot(n, dft_times, color = 'orange', label = "DFT")
ax1.legend()
ax1.grid(True)
ax1.legend()
ax1.set_title("Comparação entre DFT e FFT")
ax1.set_ylabel("Tempo de Execução (segundos)")


ax2.plot(n, fft_times, color = 'green', label = "FFT")
ax2.plot(n, dft_times, color = 'orange', label = "DFT")
ax2.legend()
ax2.grid(True)
ax2.legend()
ax2.set_yscale('log')
ax2.set_xlabel("Grau do Polinômio")
ax2.set_ylabel("Tempo de Execução (log-segundos)")

plt.savefig(data_path.parent / 'execution-comparison.png', dpi=1000, bbox_inches='tight')

# print(n)
# print(dft_times)
# print(fft_times)