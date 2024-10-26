import json
import matplotlib.pyplot as plt
import numpy as np

file_path = "../golang/count-sort/comparison-counting-and-radix.json"
with open(file_path, "r") as file:
    data = json.load(file)

how_much = len(data['n']) // 2
print(data.keys())
MAX = 2e+07

n = np.array(data['n'])
radix = np.array(data['tempos_radix'])
counting = np.array(data['tempos_counting'])

mask_max = (counting < MAX)*(radix < MAX)
mask_zero = (counting > 0)*(radix > 0)
mask = mask_max*mask_zero

n = n[mask]
radix = radix[mask]
counting = counting[mask]

plt.scatter(n, radix, color='red',s=8, label='Radix Sort')

plt.scatter(n, counting, color='blue',s=8, label='Counting sort')

plt.title("Sorting Algorithms Running Time")
plt.xlabel("Array length")
plt.ylabel("Running Time (ns)")

plt.grid(True)
plt.legend()
plt.savefig(
    "../golang/count-sort/comparison.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()