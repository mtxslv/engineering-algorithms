import re
from pathlib import Path
import matplotlib.pyplot as plt
import numpy as np

logs = [
    "Took 4732 nanoseconds to insert in Hash Table",
    "Took 5466 nanoseconds to insert with Open Addressing",
    "Took 4385 nanoseconds to search on Hash Table",
    "Took 4767 nanoseconds to search on Open Addressing",
    "Took 4892 nanoseconds to insert in Hash Table",
    "Took 5699 nanoseconds to insert with Open Addressing",
    "Took 2907 nanoseconds to search on Hash Table",
    "Took 4598 nanoseconds to search on Open Addressing",
    "Took 5138 nanoseconds to insert in Hash Table",
    "Took 5846 nanoseconds to insert with Open Addressing",
    "Took 2903 nanoseconds to search on Hash Table",
    "Took 4722 nanoseconds to search on Open Addressing",
    "Took 5096 nanoseconds to insert in Hash Table",
    "Took 5836 nanoseconds to insert with Open Addressing",
    "Took 2876 nanoseconds to search on Hash Table",
    "Took 4631 nanoseconds to search on Open Addressing",
    "Took 7248 nanoseconds to insert in Hash Table",
    "Took 8158 nanoseconds to insert with Open Addressing",
    "Took 3657 nanoseconds to search on Hash Table",
    "Took 5586 nanoseconds to search on Open Addressing",
    "Took 4912 nanoseconds to insert in Hash Table",
    "Took 5957 nanoseconds to insert with Open Addressing",
    "Took 2840 nanoseconds to search on Hash Table",
    "Took 4756 nanoseconds to search on Open Addressing"
]
pattern = r'Took (\d.+) nanoseconds to (insert|search) (on|with|in) (Hash Table|Open Addressing)'
data = {
    'insert': {
        'Hash Table': [],
        'Open Addressing': []
    },
    'search': {
        'Hash Table': [],
        'Open Addressing': []
    }
}
for log in logs:
    # ('4722', 'search', 'on', 'Open Addressing')
    info = re.findall(pattern, log)[0]
    data[info[1]][info[-1]].append(int(info[0]))

# Create subplots
fig, axs = plt.subplots(2, 1, figsize=(10, 8))

# Plot the first subgraph (insert)
x = np.arange(len(data['insert']['Hash Table']))
width = 0.35
axs[0].bar(x - width/2, data['insert']['Hash Table'],
           width, label='Tabela Hash')
axs[0].bar(x + width/2, data['insert']['Open Addressing'],
           width, label='Endereçamento Aberto')
axs[0].set_ylabel('Tempo de Execução (ns)')
axs[0].set_title('Adição')
axs[0].set_xticks(x)
axs[0].set_xticklabels(['Experimento  1', 'Experimento  2', 'Experimento  3',
                       'Experimento  4', 'Experimento  5', 'Experimento  6'])
axs[0].legend()

# Plot the second subgraph (search)
x = np.arange(len(data['search']['Hash Table']))
axs[1].bar(x - width/2, data['search']['Hash Table'],
           width, label='Tabela Hash')
axs[1].bar(x + width/2, data['search']['Open Addressing'],
           width, label='Endereçamento Aberto')
axs[1].set_ylabel('Tempo de Execução (ns)')
axs[1].set_title('Busca')
axs[1].set_xticks(x)
axs[1].set_xticklabels(['Experimento  1', 'Experimento  2', 'Experimento  3',
                       'Experimento  4', 'Experimento  5', 'Experimento  6'])
axs[1].legend()

plt.tight_layout()
plt.savefig(
    Path().cwd() / 'experiments' / "fruit-hash" / "fruits-hash.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()
