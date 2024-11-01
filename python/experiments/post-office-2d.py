import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

# Define the function
def post_office(x, y):
    points_x = [1, 2, 3, 4, 5, 6, 8]
    points_y = [-3, -2, 0, 4, 5, 7, 18]
    weights = [0.075, 0.025, 0.12, 0.15, 0.08, 0.2, 0.35]
    ans = 0
    for it in range(len(weights)):
        ans += weights[it] * (np.abs(x - points_x[it]) + np.abs(y - points_y[it]))
    return ans

# Create a grid
x_values = np.arange(0, 10, 0.01)
y_values = np.arange(-5, 20, 0.01)
X, Y = np.meshgrid(x_values, y_values)

# Calculate z values for each (x, y) pair
Z = np.vectorize(post_office)(X, Y)

# Create subplots
fig = plt.figure(figsize=(14, 6))

# Left subplot: 3D surface plot
ax1 = fig.add_subplot(1, 2, 1, projection='3d')
ax1.plot_surface(X, Y, Z, cmap='viridis')
ax1.set_xlabel('X')
ax1.set_ylabel('Y')
ax1.set_zlabel('Weighted Distance')
ax1.set_title('3D Surface of Weighted Distance')

# Right subplot: 2D color map
ax2 = fig.add_subplot(1, 2, 2)
contour = ax2.contourf(X, Y, Z, cmap='viridis')
ax2.axvline(6, color='orange', linestyle='dashed',label='weighted median')
ax2.axhline(7, color='orange', linestyle='dashed',label='weighted median')
fig.colorbar(contour, ax=ax2, label='Weighted Distance')
ax2.set_xlabel('X')
ax2.set_ylabel('Y')
ax2.set_title('2D Contour Map of Weighted Distance')

plt.tight_layout()
plt.show()
