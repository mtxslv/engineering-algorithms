from pathlib import Path
import re
from matplotlib import pyplot as plt

data_path = Path(__file__).parents[2] / 'golang' / 'LRU-and-random-marking' / 'results.txt'
assert data_path.exists() and data_path.is_file()

with open(data_path) as file:
    file_contents = file.readlines()

pattern = r"Total Requests \(N\): 336 \| Cache Size \(K\): 15 \| LRU Misses: (\d.+) \| RMC Misses: (\d.+) \| Competitiveness \(LRU\/OPT\): (.+)"

all_misses_for_lru = []
all_misses_for_rmc = []
trials = []
all_c = []

for it, line in enumerate(file_contents):
    lru_misses, rmc_misses, competitiveness = re.findall(pattern, line)[0]
    lru_misses, rmc_misses, competitiveness = int(lru_misses), int(rmc_misses), float(competitiveness)
    all_misses_for_lru.append(lru_misses)
    all_misses_for_rmc.append(rmc_misses)
    all_c.append(competitiveness)
    trials.append(it+1)

fig, (ax1,ax2) = plt.subplots(2,1, figsize=(10, 8), sharex=True)
ax1.plot(trials,all_misses_for_lru,color='green',label="LRU miss")
ax1.plot(trials,all_misses_for_rmc,color='orange',label="RMC miss")
ax1.legend()
ax1.grid(True)

ax2.plot(trials, all_c, label="Competitiveness (LRU/RMC)")
ax2.legend()
ax2.grid(True)

ax1.set_title("Performance Memórias Cache")
plt.xlabel("Experimentos")
plt.savefig(data_path.parent / 'experiments.png', dpi=1000, bbox_inches='tight')
