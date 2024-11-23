from pathlib import Path

import numpy as np
import osmnx as ox


class Vertex:
    parent = None
    d = np.inf # upper bound

def load_graph():
    file_path = Path().home() / "Documentos/engineering-algorithms/python/graph-extraction/recife_praca_comunidade.graphml"
    return ox.load_graphml(str(file_path))

def main():
    SOURCE_NODE_ID = 3691433990
    TARGET_NODE_ID = 3921998309
    G = load_graph()
    print(G.out_edges(SOURCE_NODE_ID,data=True))
    
if __name__=='__main__':
    main()