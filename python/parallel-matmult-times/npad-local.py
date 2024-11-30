import matplotlib.pyplot as plt
import numpy as np
from pathlib import Path
# Data
data = {
    'local': [505930680, 514085419, 505939174],
    'npad': [276781478, 279281763, 277822766]
}

# Categories
categories = list(data.keys())
values = [data['local'], data['npad']]

# Create index for each group
indices = np.arange(len(data['local']))

# Bar width
bar_width = 0.35

# Plot bars
plt.bar(indices, data['local'], width=bar_width, label='Local', color='limegreen')
plt.bar(indices + bar_width, data['npad'], width=bar_width, label='NPad', color='royalblue')

# Add labels and title
plt.xlabel('Experimentos')
plt.ylabel('Tempo (microsegundos)')
plt.title('Processamento de Matrizes')
plt.xticks(indices + bar_width / 2, ["1", "2", "3"])
plt.legend()

# Display plot
plt.tight_layout()
plt.savefig(
    Path().cwd() /  "parallel-matmult-times" / "npad-local.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()
