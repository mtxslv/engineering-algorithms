import osmnx as ox

distance = 1550
city_address = 'Recife, Pernambuco'
pt = (-8.056556,-34.8731439)
graph = ox.graph_from_point(pt, dist=distance)

# Plot the graph
ox.plot_graph(graph)

# Save the graph to a file (optional)
ox.save_graphml(graph, "recife_praca_comunidade.graphml")

print("Street graph extracted and saved!")
