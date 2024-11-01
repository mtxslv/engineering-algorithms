import numpy as np
import matplotlib.pyplot as plt

points = [1, 2, 3, 4, 5, 6, 8]
weights = [0.075, 0.025, 0.12, 0.15, 0.08, 0.2, 0.35]

def post_office(x):
    assert len(points) == len(weights)
    y = 0
    for it, pt in enumerate(points):
        y += np.abs(x-pt)*weights[it]
    return y

x_values = np.arange(0,10,0.01)
y_values = map(
    post_office,
    x_values
)

plt.axvline(6, color='orange', linestyle='dashed',label='weighted median')
plt.plot(x_values,list(y_values))
plt.legend()
plt.show()