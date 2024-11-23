import numpy as np
from pathlib import Path

import matplotlib.pyplot as plt
import osmnx as ox

file_path = Path().home() / "Documentos/engineering-algorithms/python/graph-extraction/recife_praca_comunidade.graphml"
print(file_path)
G = ox.load_graphml(str(file_path))

nodes = list(G.nodes(data=True))
edges = list(G.edges(data=True))
nodes_amorim = set()
nodes_museu_militar = set()
for edge in edges:
    if 'name' not in edge[2]:
        continue
    else:
        name = edge[2]['name']
        if isinstance(name,str):
            if 'Amorim' in name:
                nodes_amorim.add(edge[0])
                nodes_amorim.add(edge[1])
            elif 'Praça da Comunidade Luso Brasileira' in name:
                nodes_museu_militar.add(edge[0])
                nodes_museu_militar.add(edge[1])

print(nodes_amorim)
print(nodes_museu_militar)

print(G.nodes[3691433988])
##############################################################################3
def manhattan_distance(target_point, test_point):
    target = np.array(target_point)
    test = np.array(test_point)
    dist = np.sum(np.abs(target-test))  
    # print(dist)
    return dist

# TRAVESSA AMORIM
target_point = (-8.065389744332165, -34.87311026037246)
smallest_dist = np.inf
closest_node = None
closest_dist = np.inf
for node_id in nodes_amorim:
    node = G.nodes[node_id]
    test_point = (node['y'],node['x'])
    dist = manhattan_distance(target_point, test_point)
    if dist < smallest_dist:
        closest_node = node
        closes_node_id = node_id
        closest_dist = dist

print(f'The closest node is {closest_node} (id: {closes_node_id})\nClosest dist: {closest_dist}')

print("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")

target_point = -8.053184433002896, -34.87154117182091
smallest_dist = np.inf
closest_node = None
closest_dist = np.inf
for node_id in nodes_museu_militar:
    node = G.nodes[node_id]
    test_point = (node['y'],node['x'])
    dist = manhattan_distance(target_point, test_point)
    if dist < smallest_dist:
        closest_node = node
        closes_node_id = node_id
        closest_dist = dist

print(f'The closest node is {closest_node} (id: {closes_node_id})\nClosest dist: {closest_dist}')

# ox.plot_graph(G)

# A ideia é ir do Mirante do Paço, que fica na Tv. do Amorim próximo a coordenada
#  (-8.065389744332165, -34.87311026037246)
# MIRANTE-DO-PAÇO É O NÓ 3691433990
# até o Museu Militar do Forte do Brum, que fica na Av. Militar, próximo a coordenada
#  (-8.053480397436534, -34.8714500267351)
# MUSEU-MILITAR-DO-FORTE-DO-BRUM É O NÓ 3921998309
# 