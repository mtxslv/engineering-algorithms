import json
from pathlib import Path

import matplotlib.pyplot as plt
import numpy as np

# Load JSON data from files
with open('results/annotated-insertion-sort/times.json', 'r') as file:
    data_times = json.load(file)

with open('results/annotated-insertion-sort/ops.json', 'r') as file:
    data_ops = json.load(file)

# Convert keys to x-values (as integers) for both datasets
x_values_times = list(map(int, data_times.keys()))
x_values_ops = list(map(int, data_ops.keys()))

# Compute mean and standard deviation for times data
means_times = [np.mean(data_times[key]) for key in data_times]
std_devs_times = [np.std(data_times[key]) for key in data_times]

# Compute mean and standard deviation for ops data
means_ops = [np.mean(data_ops[key]) for key in data_ops]
std_devs_ops = [np.std(data_ops[key]) for key in data_ops]

# Create figure and axis for the first plot (times data)
fig, ax1 = plt.subplots()

# Plot the average (mean) curve for times data
ax1.plot(x_values_times, means_times, label='Mean times curve', color='blue')

# Add the painted region for the standard deviation (times data)
ax1.fill_between(x_values_times, 
                 np.array(means_times) - np.array(std_devs_times), 
                 np.array(means_times) + np.array(std_devs_times), 
                 color='blue', alpha=0.2, label='Times std dev region')

# Labels for the first plot (times data)
ax1.set_xlabel('List lengths')
ax1.set_ylabel('Times (ns)', color='blue')
ax1.set_title('Insertion Sort Performance')

# Add grid
ax1.grid(True)

# Create a second y-axis for the operations data
ax2 = ax1.twinx()

# Plot the average (mean) curve for ops data
ax2.plot(x_values_ops, means_ops, label='Mean ops curve', color='green')

# Add the painted region for the standard deviation (ops data)
ax2.fill_between(x_values_ops, 
                 np.array(means_ops) - np.array(std_devs_ops), 
                 np.array(means_ops) + np.array(std_devs_ops), 
                 color='green', alpha=0.2, label='Ops std dev region')

# Labels for the second plot (operations data)
ax2.set_ylabel('Operations', color='green')

# Add legends for both axes
ax1.legend(loc='upper left')
ax2.legend(loc='upper right')

# Save the plot 
# Save the figure
plt.savefig(
    Path().cwd() / 'results' / "insertion-performance.png",
    bbox_inches='tight',
    dpi=500
)

# Show the plot
plt.show()
