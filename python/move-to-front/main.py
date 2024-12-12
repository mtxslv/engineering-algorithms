from pathlib import Path
import re

from matplotlib import pyplot as plt

data_path = Path(__file__).parents[2] / 'golang' / 'move-to-front' / 'experiments.txt'
assert data_path.exists() and data_path.is_file()

with open(data_path) as file:
    file_contents = file.readlines()

pattern = r'LIST SIZE: (\d.+) . REQUESTS SIZE: (\d.+) .RATIO: (.+)'

data = {}
for it, exp_str in enumerate(file_contents):
    list_len, request_num, ratio = re.findall(pattern, exp_str)[0]
    list_len, request_num, ratio = int(list_len), int(request_num), float(ratio)
    data[list_len] = ratio

# Plot each list as a line
# for key, value in data.items():
x,y = list(data.keys()), list(data.values())
plt.plot(x,y, color='midnightblue')

plt.scatter(x,y, color='midnightblue')

# Add horizontal line at y=4
plt.axhline(y=4, color='black', linestyle='--', label='Competitividade')

# Set x and y axis labels
plt.xlabel('Tamanho de Lista/Requisições')
plt.ylabel('Razão entre Operações Realizadas')

# Add legend and title
plt.legend()
plt.title('Razão de Operações por Buscas')

# Grid
plt.grid(True)

# Save the figure with tight layout and high resolution
plt.savefig(data_path.parent / 'ratio.png', dpi=1000, bbox_inches='tight')

# Show the plot
plt.show()