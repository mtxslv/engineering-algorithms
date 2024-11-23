from copy import deepcopy
from pathlib import Path

import numpy as np
import osmnx as ox

class MinPriorityQueue:

    def __init__(self, main_key) -> None:
        self.queue = []  # list of dicts
        self.main_key = main_key

    def parent(self, i):
        return (i - 1) // 2

    def left(self, i):
        return 2 * i + 1

    def right(self, i):
        return 2 * i + 2

    def minimum(self):
        if len(self.queue) < 1:
            raise ValueError("Heap underflow")
        return self.queue[0]

    def min_heapify(self, i):
        """Maintain the min-heap property starting from index i."""
        l = self.left(i)
        r = self.right(i)
        smallest = i

        if l < len(self.queue) and self.queue[l][self.main_key] < self.queue[smallest][self.main_key]:
            smallest = l
        if r < len(self.queue) and self.queue[r][self.main_key] < self.queue[smallest][self.main_key]:
            smallest = r

        if smallest != i:
            self.queue[i], self.queue[smallest] = self.queue[smallest], self.queue[i]
            self.min_heapify(smallest)

    def build(self, A):
        self.queue = A
        n = len(self.queue)
        for i in range(n // 2 - 1, -1, -1):  # Start from the last non-leaf node
            self.min_heapify(i)

    def extract_min(self):
        """Remove and return the minimum element from the queue."""
        min_ = self.minimum()
        self.queue[0] = self.queue[-1]
        self.queue.pop()
        if len(self.queue) > 0:
            self.min_heapify(0)
        return min_

    def add(self, x):
        """Insert a new element x into the queue."""
        k = x[self.main_key]
        x[self.main_key] = float('inf')  # Temporarily set the key to infinity
        self.queue.append(x)
        self.decrease_key(x, k)

    def decrease_key(self, x, new_key):
        """Decrease the key of element x to new_key."""
        i = self.queue.index(x)
        if new_key > self.queue[i][self.main_key]:
            raise ValueError("New key is larger than current key")

        self.queue[i][self.main_key] = new_key

        # Fix the min-heap property by moving the element up
        while i > 0 and self.queue[self.parent(i)][self.main_key] > self.queue[i][self.main_key]:
            parent = self.parent(i)
            self.queue[i], self.queue[parent] = self.queue[parent], self.queue[i]
            i = parent

def initialize_single_source(G, source_node_id):
    for node_ID in G.nodes:
        node = G.nodes[node_ID]
        node['d'] = np.inf
        node['parent'] = None
        node['id'] = node_ID
    G.nodes[source_node_id]['d'] = 0
    return G

def dijkstra(G, w, source_node_id):
    G = initialize_single_source(G,source_node_id)
    S = set()
    vertices = deepcopy(G.nodes)
    Q = MinPriorityQueue(main_key='d')
    for vertex in vertices.values():
        Q.add(vertex)
        print(vertex)
    # while len(Q) > 0:
        # pass
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