from copy import deepcopy
from pathlib import Path

import numpy as np
import osmnx as ox

def initialize_single_source(G, source_node_id):
    for node_ID in G.nodes:
        node = G.nodes[node_ID]
        node['d'] = np.inf
        node['parent'] = None
    G.nodes[source_node_id]['d'] = 0
    return G

def dijkstra(G, w, source_node_id):
    G = initialize_single_source(G,source_node_id)
    S = set()
    Q = deepcopy(G.nodes)
    while len(Q) > 0:
        pass
    return G # change later

def load_graph():
    file_path = Path().home() / "Documentos/engineering-algorithms/python/graph-extraction/recife_praca_comunidade.graphml"
    return ox.load_graphml(str(file_path))

def main():
    SOURCE_NODE_ID = 3691433990
    TARGET_NODE_ID = 3921998309
    G = load_graph()
    G = dijkstra(G,0,TARGET_NODE_ID)
    print(G.nodes[TARGET_NODE_ID])
    # print(G.out_edges(SOURCE_NODE_ID,data=True))

if __name__=='__main__':
    main()