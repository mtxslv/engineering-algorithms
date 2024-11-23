import osmnx as ox
from pathlib import Path
import matplotlib.pyplot as plt

# Load the graph
file_path = Path().home() / "Documentos/engineering-algorithms/python/graph-extraction/recife_praca_comunidade.graphml"
G = ox.load_graphml(str(file_path))

# Identify source and target nodes
source_node = 3691433990
target_node = 3921998309

# Create a node color map
node_colors = []
for node in G.nodes():
    if node == source_node:
        node_colors.append('yellow')
    elif node == target_node:
        node_colors.append('red')
    else:
        node_colors.append('gray')  # Default color for other nodes

# Plot the graph with color-coded nodes
fig, ax = ox.plot_graph(G, node_color=node_colors, node_size=10, edge_color='gray', edge_linewidth=0.5)


# Save the figure with tight layout and high resolution
fig.savefig(file_path.parent / 'source-and-target.png', dpi=1000, bbox_inches='tight')

# (Optional) Close the plot window (useful when running in a script)
plt.close()


# plt.tight_layout()
# fig.savefig(
#     file_path.parent / 'source-and-target.png',
#     dpi=1000
# )
# plt.show()