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

if __name__ == "__main__":
    # Define the main key for the priority queue
    main_key = 'priority'

    # Initialize the MinPriorityQueue
    pq = MinPriorityQueue(main_key)

    # Add some elements to the priority queue
    pq.add({'id': 1, 'priority': 5})
    pq.add({'id': 2, 'priority': 3})
    pq.add({'id': 4, 'priority': 10})
    pq.add({'id': 3, 'priority': 8})

    print("Queue after adding elements:")
    print(pq.queue)

    # Decrease the priority of an element
    element_to_update = {'id': 3, 'priority': 8}  # Element with id=3
    new_priority = 2
    pq.decrease_key(element_to_update, new_priority)

    print("\nQueue after decreasing the priority of id=3 to 2:")
    print(pq.queue)

    min_ = pq.extract_min()

    print("\nQueue after extracting minimum:")
    print(pq.queue)

    print('\nMinimum Extracted:')
    print(min_)
