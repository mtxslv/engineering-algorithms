from pathlib import Path
import matplotlib.pyplot as plt
data = {
    1:{ # first experiment
        1:1559801532, # no thread
        5:505930680,  # 5 threads
        10:447295364, # 10 threads
        15:455503652  # 15 threads
    },
    2:{# second experiment
        1:2656958396,# no thread
        5:514085419,# 5 threads
        10:465782654,# 10 threads
        15:456578335 # 15 threads
    },
    3:{ # third experiment
        1:1516463018,# no thread
        5:505939174,# 5 threads
        10:446159692,# 10 threads
        15:455015345 # 15 threads
    },
}

# Data
# data = {
#     1: {1: 1559801532, 5: 505930680, 10: 447295364, 15: 455503652},
#     2: {1: 2656958396, 5: 514085419, 10: 465782654, 15: 456578335},
#     3: {1: 1516463018, 5: 505939174, 10: 446159692, 15: 455015345},
# }

# Markers for each experiment
markers = ['^', 'D', 'o']

# Create scatter plot
plt.figure(figsize=(10, 6))

for trial, results in data.items():
    x = list(results.keys())
    y = list(results.values())
    plt.scatter(x, y, label=f'Experimento {trial}', marker=markers[trial - 1])

# Customizing the plot
plt.title("Tempos de Execução")
plt.xlabel("Número de Threads")
plt.ylabel("Tempo (Microsegundos)")
plt.xticks([1, 5, 10, 15])
plt.legend()
plt.grid(True, which='both', linestyle='--', linewidth=0.5)
plt.tight_layout()

# Show plot
plt.tight_layout()
plt.savefig(
    Path().cwd() /  "parallel-matmult-times" / "local-times.png",
    bbox_inches='tight',
    dpi=500
)
plt.show()
