import re
from pathlib import Path
import matplotlib.pyplot as plt
import numpy as np

logs = [
    # "Took 9582 nanoseconds to search with Linear Probing",
    # "Took 6153 nanoseconds to search with Double Hashing",
    "Took 6941 nanoseconds to insert with Linear Probing",
    "Took 5957 nanoseconds to insert with Double Hashing",
    "Took 8114 nanoseconds to search with Linear Probing",
    "Took 4500 nanoseconds to search with Double Hashing",
    "Took 6519 nanoseconds to insert with Linear Probing",
    "Took 5707 nanoseconds to insert with Double Hashing",
    "Took 5687 nanoseconds to search with Linear Probing",
    "Took 4411 nanoseconds to search with Double Hashing",
    "Took 7008 nanoseconds to insert with Linear Probing",
    "Took 5943 nanoseconds to insert with Double Hashing",
    "Took 7618 nanoseconds to search with Linear Probing",
    "Took 4998 nanoseconds to search with Double Hashing",
    "Took 9625 nanoseconds to insert with Linear Probing",
    "Took 8810 nanoseconds to insert with Double Hashing",
    "Took 7321 nanoseconds to search with Linear Probing",
    "Took 5672 nanoseconds to search with Double Hashing",
    "Took 6931 nanoseconds to insert with Linear Probing",
    "Took 6108 nanoseconds to insert with Double Hashing",
    "Took 5858 nanoseconds to search with Linear Probing",
    "Took 4446 nanoseconds to search with Double Hashing"
]

pattern = r'Took (\d.+) nanoseconds to (insert|search) (on|with|in) (Double Hashing|Linear Probing)'
data = {
    'insert': {
        'Double Hashing': [],
        'Linear Probing': []
    },
    'search': {
        'Double Hashing': [],
        'Linear Probing': []
    }
}
for log in logs:
    # ('4722', 'search', 'on', 'Open Addressing')
    info = re.findall(pattern, log)[0]
    data[info[1]][info[-1]].append(int(info[0]))

print(data)
# Create subplots
fig, axs = plt.subplots(2, 1, figsize=(10, 8))

# Plot the first subgraph (insert)
x = np.arange(len(data['insert']['Double Hashing']))
width = 0.35
axs[0].bar(x - width/2, data['insert']['Double Hashing'],
           width, label='Double Hashing')
axs[0].bar(x + width/2, data['insert']['Linear Probing'],
           width, label='Linear Probing')
axs[0].set_ylabel('Tempo de Execução (ns)')
axs[0].set_title('Adição')
axs[0].set_xticks(x)
axs[0].set_xticklabels(['Experimento  1', 'Experimento  2', 'Experimento  3',
                       'Experimento  4', 'Experimento  5'])
axs[0].legend()

# Plot the second subgraph (search)
x = np.arange(len(data['search']['Double Hashing']))
axs[1].bar(x - width/2, data['search']['Double Hashing'],
           width, label='Double Hashing')
axs[1].bar(x + width/2, data['search']['Linear Probing'],
           width, label='Linear Probing')
axs[1].set_ylabel('Tempo de Execução (ns)')
axs[1].set_title('Busca')
axs[1].set_xticks(x)
axs[1].set_xticklabels(['Experimento  1', 'Experimento  2', 'Experimento  3',
                       'Experimento  4', 'Experimento  5'])
axs[1].legend()

plt.tight_layout()
plt.savefig(
    Path().cwd() / 'experiments' / "fruit-hash-linear-probing" / "fruits-hash-linear-probing.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()
