import json
import matplotlib.pyplot as plt
import numpy as np

file_path = "../golang/min-max/comparison-minmax.json"
with open(file_path, "r") as file:
    data = json.load(file)

how_much = len(data['n']) // 2
print(data.keys())
MAX = 2e+07

n = np.array(data['n'])
tempos_simultaneous = np.array(data['tempos_simultaneous'])
tempos_standalone = np.array(data['tempos_standalone'])

# mask_max = (tempos_standalone < MAX)*(tempos_simultaneous < MAX)
mask_zero = (tempos_standalone > 0)*(tempos_simultaneous > 0)
mask = mask_zero # mask_max*

n = n[mask]
tempos_simultaneous = tempos_simultaneous[mask]
tempos_standalone = tempos_standalone[mask]

plt.scatter(n, tempos_simultaneous, color='red',s=8, label='Simultaneous', alpha=0.3)

plt.scatter(n, tempos_standalone, color='blue',s=8, label='Standalone', alpha=0.3)

plt.title("Sorting Algorithms Running Time")
plt.xlabel("Array length")
plt.ylabel("Running Time (ns)")

plt.grid(True)
plt.legend()
plt.savefig(
    "../golang/min-max/comparison.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()