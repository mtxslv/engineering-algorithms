import osmnx as ox
from shapely.geometry import Polygon

# Define the points for the polygon
pts = [
    (-8.042879587934328, -34.86956669792991),
    (-8.047342031608311, -34.87136934353605),
    (-8.051294229898248, -34.87349456727075),
    (-8.063116131320255, -34.87541615138023),
    (-8.069147258723254, -34.87296250290411),
    (-8.049629190267684, -34.86389487765625),
    (-8.044788584489817, -34.8643680387973),
    (-8.042879587934328, -34.86956669792991),
]

# Create the polygon
polygon = Polygon(pts)

# Define tags to filter OSM features (optional, for roads we often use {"highway": True})
# tags = {"highway": True}

# Retrieve OSM features inside the polygon
# features = ox.features_from_polygon(polygon, tags)

# Extract the street graph within the polygon
graph = ox.graph_from_polygon(polygon)

# Plot the graph
ox.plot_graph(graph)

# Optionally save the graph
ox.save_graphml(graph, "recife_polygon_region.graphml")

print("Graph extracted from the polygon and saved!")
