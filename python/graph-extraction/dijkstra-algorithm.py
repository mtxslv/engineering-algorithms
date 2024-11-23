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

    def __len__(self):
        return len(self.queue)
    
    def pop_node_by_id(self, node_id):
        """
        Since it is impossible to decrease elements in heap using float
        let's just pop old node and add it back in the main code. 

        Parameters
        ----------
        node_id : _type_
            _description_
        """
        pos = -1
        for it, node in enumerate(self.queue):
            if node['id'] == node_id:
                pos = it
                break
        if pos > 0:
            del self.queue[pos]

def initialize_single_source(G, source_node_id):
    for node_ID in G.nodes:
        node = G.nodes[node_ID]
        node['d'] = np.inf
        node['parent'] = None
        node['id'] = node_ID
    G.nodes[source_node_id]['d'] = 0
    return G

def relax(node_u, node_v, weight_from_u_to_v):
    if node_v['d'] > node_u['d'] + weight_from_u_to_v:
        node_v['d'] = node_u['d'] + weight_from_u_to_v
        node_v['parent'] = node_u['id']
        # print(node_v['id'])
        return True
    else:
        return False

def dijkstra(G, source_node_id): 
    # No weight function w is supplied, given 
    # the graph edges are weighted themselves. 
    G = initialize_single_source(G,source_node_id)
    S = []
    vertices = deepcopy(G.nodes)
    Q = MinPriorityQueue(main_key='d')
    for vertex in vertices.values():
        Q.add(vertex)
    while len(Q) > 0:
        print(f"QUEUE SIZE: {len(Q)}")
        u = Q.extract_min()
        S.append(u)
        for edge_info in G.out_edges(u['id'],data=True):
            # edge_info[1] is adj vertex index
            v = G.nodes[edge_info[1]]  
            weight_u_to_v = edge_info[-1]['length']
            decreased_node_v = relax(u,v,weight_u_to_v)
            if decreased_node_v:
                # The following three lines tries to ammend 
                # the priority queue, once decrease key
                # does not support float values for keys.
                Q.pop_node_by_id(v['id'])
                Q.min_heapify(len(Q)//2-1)
                Q.add(v)

    return S 

def load_graph():
    file_path = Path().home() / "Documentos/engineering-algorithms/python/graph-extraction/recife_praca_comunidade.graphml"
    return ox.load_graphml(str(file_path))

def plot_path(G, path, source_node, target_node):
    # Create a node color map
    node_colors = []
    for node in G.nodes():
        if node == source_node:
            node_colors.append('yellow')
        elif node == target_node:
            node_colors.append('red')
        elif node in path:
            node_colors.append('orange')
        else:
            node_colors.append('gray')  # Default color for other nodes

    # Plot the graph with color-coded nodes
    fig, ax = ox.plot_graph(G, node_color=node_colors, node_size=10, edge_color='gray', edge_linewidth=0.5)

    file_path = Path().home() / "Documentos/engineering-algorithms/python/graph-extraction/"
    fig.savefig(file_path / 'path.png', dpi=1000, bbox_inches='tight')

def main():
    SOURCE_NODE_ID = 3691433990
    TARGET_NODE_ID = 3921998309
    G = load_graph()
    S = dijkstra(G,SOURCE_NODE_ID)

    path = []
    parent_node_id = G.nodes[TARGET_NODE_ID]['parent']
    while parent_node_id != SOURCE_NODE_ID:
        path.append(parent_node_id)
        parent_node_id = G.nodes[parent_node_id]['parent']

    plot_path(G,path,SOURCE_NODE_ID,TARGET_NODE_ID)    

if __name__=='__main__':
    main()