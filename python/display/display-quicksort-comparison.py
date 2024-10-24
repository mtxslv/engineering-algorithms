import json
import matplotlib.pyplot as plt
import numpy as np

file_path = "../golang/quicksort/comparison-quicksort-and-randomized.json"
with open(file_path, "r") as file:
    data = json.load(file)

how_much = len(data['n']) // 2
print(data.keys())
MAX = 2e+07

n = np.array(data['n'])
randomized = np.array(data['tempos_randomized'])
regular = np.array(data['tempos_regular'])

mask_max = (regular < MAX)*(randomized < MAX)
mask_zero = (regular > 0)*(randomized > 0)
mask = mask_max*mask_zero

n = n[mask]
randomized = randomized[mask]
regular = regular[mask]

plt.scatter(n, randomized, color='red',s=8, label='Randomized-QuickSort')

plt.scatter(n, regular, color='blue',s=8, label='Quicksort')

plt.title("Sorting Algorithms Running Time")
plt.xlabel("Array length")
plt.ylabel("Running Time (ns)")

plt.grid(True)
plt.legend()
plt.savefig(
    "../golang/quicksort/comparison.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()