from pathlib import Path
import re
import numpy as np
from matplotlib import pyplot as plt
import plotly.graph_objects as go
import numpy as np
import pandas as pd
data_path = Path(__file__).parents[2] / 'golang' / 'least-recently-used' / 'ans.txt'
assert data_path.exists() and data_path.is_file()

with open(data_path) as file:
    file_contents = file.readlines()


pattern = r"Total Requests \(N\): (\d.+) \| Cache Size \(K\): (\d+) \| LRU Misses: (\d.+) \| OPT Misses: (\d.+) \| Competitiveness \(LRU\/OPT\): (.+)"

data = {'N':[],'K':[],'C':[]}
for line in file_contents:
    n,k, lru_misses, opt_misses,competitiveness = re.findall(pattern, line)[0]
    n,k = int(n), int(k)
    competitiveness = float(competitiveness)
    # print(f"N={n}, K={k} || {lru_misses}/{opt_misses} = {competitiveness}")
    data['N'].append(n)
    data['K'].append(k)
    data['C'].append(competitiveness)
# Convert the dictionary to a pandas DataFrame for easier handling
df = pd.DataFrame(data)

# Create the 3D scatter plot
fig = go.Figure(data=[go.Scatter3d(
    x=df['N'],  # x-axis: N
    y=df['K'],  # y-axis: K
    z=df['C'],  # z-axis: C
    mode='markers',  # Show markers (points)
    marker=dict(
        size=8,  # Adjust marker size
        color=df['C'],  # Color markers based on C values
        colorscale='Viridis',  # Choose a colorscale
        opacity=0.8  # Adjust marker opacity
    )
)])

# Customize the layout (optional)
fig.update_layout(
    scene=dict(
        xaxis_title='N',
        yaxis_title='K',
        zaxis_title='C',
        aspectmode='cube' #important to avoid stretching
    ),
    margin=dict(l=0, r=0, b=0, t=0) #remove margins for better view
)

# fig.show()
fig.write_html(Path(__file__).parents[2] / 'golang' / 'least-recently-used' / "results.html")
