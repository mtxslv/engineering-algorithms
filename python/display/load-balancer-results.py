import json
import matplotlib.pyplot as plt
from pathlib import Path
# Load the JSON data from file
with open('load-balancer-problem/loads.json', 'r') as file:
    data = json.load(file)

# Get the sorted task numbers (outermost keys)
n_values = sorted(map(int, data.keys()))

# Prepare lists for each worker's hours (original and shuffled)
original_workers = [[] for _ in range(4)]
shuffled_workers = [[] for _ in range(4)]

# Extract hours worked by each worker
for n in n_values:
    for i in range(4):
        original_workers[i].append(data[str(n)]["original"][i])
        shuffled_workers[i].append(data[str(n)]["shuffled"][i])

# Colors for original workers
original_colors = ['slategrey', 'darkslateblue', 'steelblue', 'navy']
# Colors for shuffled workers
shuffled_colors = ['orangered', 'darkorange', 'orange', 'peru']

# Create the plot
plt.figure(figsize=(10, 6))

# Plot lines for the original workers with appropriate colors and labels
for i in range(4):
    plt.plot(n_values, original_workers[i], label=f'{i+1}st worker (original)' if i == 0 else f'{i+1}nd worker (original)' if i == 1 else f'{i+1}rd worker (original)' if i == 2 else f'{i+1}th worker (original)', color=original_colors[i], linestyle='-', linewidth=2)

# Plot lines for the shuffled workers with appropriate colors and labels
for i in range(4):
    plt.plot(n_values, shuffled_workers[i], label=f'{i+1}st worker (shuffled)' if i == 0 else f'{i+1}nd worker (shuffled)' if i == 1 else f'{i+1}rd worker (shuffled)' if i == 2 else f'{i+1}th worker (shuffled)', color=shuffled_colors[i], linestyle='-', linewidth=2)

# Add labels and title
plt.xlabel('Number of tasks per worker')
plt.ylabel('Load')
plt.title('Load per Worker (Original vs Shuffled Tasks)')

# Add legend
plt.legend(loc='upper left')

# Add grid
plt.grid(True)

# Save the plot 
# Save the figure
plt.savefig(
    Path().cwd() / 'results' / 'images'/ "load-balancer-results.png",
    bbox_inches='tight',
    dpi=500
)


# Show the plot
plt.show()
