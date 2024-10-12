import json
import matplotlib.pyplot as plt
from pathlib import Path

# Load the JSON data from 'comparison.json'
with open('./results/comparing-insertion-and-merge/comparison.json', 'r') as f:
    tempos_ = json.load(f)

# Set up a 2x2 subplot layout
fig, axs = plt.subplots(2, 2, figsize=(12, 10))

# Flatten the 2x2 grid of subplots for easier indexing
axs = axs.flatten()

# Define titles for each subplot and their corresponding index
subplot_titles = ['Best', 'Mid-Best', 'Mid-Worst', 'Worst']

# Loop through each key in tempos_ and plot the respective curves in its own subplot
for idx, (case, values) in enumerate(tempos_.items()):
    n_values = values['n']
    insertion_times = values['insertion']
    merge_times = values['merge']
    
    # Plot insertion sort times (solid line)
    axs[idx].plot(n_values, insertion_times, label='Insertion Sort')
    
    # Plot merge sort times (dotted line)
    axs[idx].plot(n_values, merge_times, label='Merge Sort')

    # Set title for the subplot
    axs[idx].set_title(subplot_titles[idx])
    
    # Set labels for the x and y axes
    if idx > 1:
        axs[idx].set_xlabel('Array Length (n)')
    axs[idx].set_ylabel('Time (nanoseconds)')

    # Add legend to each subplot
    axs[idx].legend()

    # Add grid for better readability
    axs[idx].grid(True)

# Adjust layout so that plots don't overlap
plt.tight_layout()

# Save the plot 
# Save the figure
plt.savefig(
    Path().cwd() / 'results' / "comparison-insertion-merge.png",
    bbox_inches='tight',
    dpi=500
)

# Show the plot
plt.show()
