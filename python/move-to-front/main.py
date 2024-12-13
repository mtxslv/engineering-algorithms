from pathlib import Path
import re
import numpy as np
from matplotlib import pyplot as plt

data_path = Path(__file__).parents[2] / 'golang' / 'move-to-front' / 'experiments.txt'
assert data_path.exists() and data_path.is_file()

with open(data_path) as file:
    file_contents = file.readlines()

pattern = r'(DIFF|EQUAL) | LIST SIZE: (\d.+) . REQUESTS SIZE: (\d.+) .RATIO: (.+)'

data = {
    'DIFF': {},
    'EQUAL': {}
}
for it, exp_str in enumerate(file_contents):
    found = re.findall(pattern, exp_str)
    operation = found[0][0]
    _, list_len, request_num, ratio = found[1]
    list_len, request_num, ratio = int(list_len), int(request_num), float(ratio)
    data[operation][request_num] = ratio

fig, axes = plt.subplots(2, 1, figsize=(10, 8))

# First subplot (EQUAL)
x, y = list(data['EQUAL'].keys()), list(data['EQUAL'].values())
axes[0].plot(x, y, color='midnightblue')
axes[0].scatter(x, y, color='midnightblue', label='Pior Caso')
axes[0].axhline(y=4, color='black', linestyle='--', label='Competitividade')
# axes[0].set_xlabel('Tamanho de Lista/Requisições')
axes[0].set_ylabel('Razão entre Operações Realizadas')
axes[0].legend()
axes[0].set_title('Razão de Operações por Buscas')
axes[0].grid(True)

# Second subplot (DIFF)
x, y = list(data['DIFF'].keys()), list(data['DIFF'].values())
axes[1].plot(x, y, color='midnightblue')
axes[1].scatter(x, y, color='midnightblue', label="Aleatório")
axes[1].set_xlabel('Número de Requisições')
axes[1].set_ylabel('Razão entre Operações Realizadas')
axes[1].legend()
axes[1].grid(True)

def tick_function(x):
    v = x/900
    return ["%.1f" % z for z in v]

ax2 = axes[1].twiny()
new_tick_locations = np.array([18e3,36e3, 54e3, 72e3, 90e3])
ax2.set_xticks(new_tick_locations)
ax2.set_xticklabels(tick_function(new_tick_locations))

# Save the figure with tight layout and high resolution
plt.savefig(data_path.parent / 'experiments.png', dpi=1000, bbox_inches='tight')

# # Show the plot
# plt.show()