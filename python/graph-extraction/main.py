import osmnx as ox

# Define the bounding box for Praça da Comunidade Luso Brasileira
north, south, east, west = -8.0605, -8.0625, -34.8728, -34.8748

# Download and extract the street graph for the defined area
graph = ox.graph_from_bbox(north=north, south=south, east=east, west=west, network_type="drive")

# Plot the graph
ox.plot_graph(graph)

# Save the graph to a file (optional)
ox.save_graphml(graph, "recife_praca_comunidade.graphml")

print("Street graph extracted and saved!")
