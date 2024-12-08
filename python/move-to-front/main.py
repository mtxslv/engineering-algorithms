from pathlib import Path
import re

from matplotlib import pyplot as plt

data_path = Path(__file__).parents[2] / 'golang' / 'move-to-front' / 'experiments.txt'
assert data_path.exists() and data_path.is_file()

with open(data_path) as file:
    file_contents = file.readlines()

contents_str = ''.join(file_contents)
experiments_str = contents_str.split('\n\n\n')[1:]

assert len(experiments_str) == 5

pattern = r'\(ratio = (\d\.\d+)\)'

data = {}
for it, exp_str in enumerate(experiments_str):
    chunk = re.findall(pattern, exp_str)
    values = [float(c) for c in chunk]
    assert len(values) == 50
    data[it+1] = values

# Plot each list as a line
for key, value in data.items():
    plt.plot(value, label=f"Experimento #{key}")

# Add horizontal line at y=4
plt.axhline(y=4, color='black', linestyle='--', label='Competitividade')

# Set x and y axis labels
plt.xlabel('Buscas Efetuadas')
plt.ylabel('Razão entre Operações Realizadas')

# Add legend and title
plt.legend()
plt.title('Razão de Operações por Música Buscada')

# Grid
plt.grid(True)

# Save the figure with tight layout and high resolution
plt.savefig(data_path.parent / 'ratio.png', dpi=1000, bbox_inches='tight')

# Show the plot
plt.show()