import json
import matplotlib.pyplot as plt
from pathlib import Path
import math
# Load the JSON data from file

with open('./results/merge-sort/result.json', 'r') as file:
    data = json.load(file)

n = data['n']
times = data['time']

constante = 107
y = [
    (constante*i)*math.log(constante*i) for i in n
]
# Create the plot
plt.figure(figsize=(10, 6))

plt.plot(n,times,linestyle='-', linewidth=2,label="Running Time")
plt.plot(n,y,linestyle='-', linewidth=2, label="107n*ln(107n)")
plt.legend()
plt.title("Merge Sort Running Time")
plt.xlabel("Array Size")
plt.grid(True)
plt.savefig(
    Path().cwd() / 'results' / "merge-sort"/"merge-sort-performance.png",
    bbox_inches='tight',
    dpi=500    
)
plt.show()