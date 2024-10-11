import json
import matplotlib.pyplot as plt

# Load the JSON data
with open('comparison.json', 'r') as f:
    data = json.load(f)

# Extract the relevant data from the dictionary
outer_key = 'worst'
n_values = data[outer_key]['n']
insertion_times = data[outer_key]['insertion']
merge_times = data[outer_key]['merge']

# Create a plot
plt.figure(figsize=(10, 6))

# Plot insertion sort times
plt.plot(n_values, insertion_times, label='Insertion Sort')

# Plot merge sort times
plt.plot(n_values, merge_times, label='Merge Sort')

# Add labels and title
plt.xlabel('n (size of the list)')
plt.ylabel('Time (microseconds)')
plt.title(f'Performance Comparison: {outer_key}')
plt.legend()

# Show the plot
plt.grid(True)
plt.show()
