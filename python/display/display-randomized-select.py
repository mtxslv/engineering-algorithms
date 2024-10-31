import json
import matplotlib.pyplot as plt
import numpy as np

file_path = "../golang/select/comparison-select-shuffled.json"
with open(file_path, "r") as file:
    data = json.load(file)

how_much = len(data['n']) // 2
print(data.keys())
MAX = 2e+07

n = np.array(data['n'])
tempos_random_select = np.array(data['tempos_random_select'])
tempos_select = np.array(data['tempos_select'])

# mask_max = (tempos_select < MAX)*(tempos_random_select < MAX)
mask_zero = (tempos_select > 0)*(tempos_random_select > 0)
mask = mask_zero # mask_max*

n = n[mask]
tempos_random_select = tempos_random_select[mask]
tempos_select = tempos_select[mask]

plt.scatter(n, tempos_random_select, color='red',s=8, label='Random Select', alpha=0.1)

plt.scatter(n, tempos_select, color='blue',s=8, label='Select', alpha=0.1)

plt.title("Search Algorithms Running Time")
plt.xlabel("Array length")
plt.ylabel("Running Time (ns)")

plt.grid(True)
plt.legend()
plt.savefig(
    "../golang/select/comparison-shuffled.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()