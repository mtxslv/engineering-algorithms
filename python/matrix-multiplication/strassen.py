def add_matrix(A, B):
    """
    Add two matrices A and B.
    """
    size = len(A)
    result = [[0] * size for _ in range(size)]
    for i in range(size):
        for j in range(size):
            result[i][j] = A[i][j] + B[i][j]
    return result

def subtract_matrix(A, B):
    """
    Subtract matrix B from matrix A.
    """
    size = len(A)
    result = [[0] * size for _ in range(size)]
    for i in range(size):
        for j in range(size):
            result[i][j] = A[i][j] - B[i][j]
    return result

def split_matrix(matrix):
    """
    Split a matrix into four submatrices (quadrants).
    """
    size = len(matrix)
    half = size // 2

    A11 = [[matrix[i][j] for j in range(half)] for i in range(half)]
    A12 = [[matrix[i][j] for j in range(half, size)] for i in range(half)]
    A21 = [[matrix[i][j] for j in range(half)] for i in range(half, size)]
    A22 = [[matrix[i][j] for j in range(half, size)] for i in range(half, size)]

    return A11, A12, A21, A22

def combine_matrix(A11, A12, A21, A22):
    """
    Combine four submatrices into one larger matrix.
    """
    size = len(A11) * 2
    result = [[0] * size for _ in range(size)]

    half = size // 2
    for i in range(half):
        for j in range(half):
            result[i][j] = A11[i][j]
            result[i][j + half] = A12[i][j]
            result[i + half][j] = A21[i][j]
            result[i + half][j + half] = A22[i][j]

    return result

def strassen(A, B):
    """
    Perform Strassen matrix multiplication on two matrices A and B.
    """
    size = len(A)

    # Base case for 1x1 matrix
    if size == 1:
        return [[A[0][0] * B[0][0]]]

    # Split the matrices into quadrants
    A11, A12, A21, A22 = split_matrix(A)
    B11, B12, B21, B22 = split_matrix(B)

    # Compute the 7 Strassen products
    M1 = strassen(add_matrix(A11, A22), add_matrix(B11, B22))    # M1 = (A11 + A22) * (B11 + B22)
    M2 = strassen(add_matrix(A21, A22), B11)                    # M2 = (A21 + A22) * B11
    M3 = strassen(A11, subtract_matrix(B12, B22))               # M3 = A11 * (B12 - B22)
    M4 = strassen(A22, subtract_matrix(B21, B11))               # M4 = A22 * (B21 - B11)
    M5 = strassen(add_matrix(A11, A12), B22)                    # M5 = (A11 + A12) * B22
    M6 = strassen(subtract_matrix(A21, A11), add_matrix(B11, B12))  # M6 = (A21 - A11) * (B11 + B12)
    M7 = strassen(subtract_matrix(A12, A22), add_matrix(B21, B22))  # M7 = (A12 - A22) * (B21 + B22)

    # Combine the results to get the final quadrants of the result matrix
    C11 = add_matrix(subtract_matrix(add_matrix(M1, M4), M5), M7)
    C12 = add_matrix(M3, M5)
    C21 = add_matrix(M2, M4)
    C22 = add_matrix(subtract_matrix(add_matrix(M1, M3), M2), M6)

    # Combine the four quadrants into a single matrix
    return combine_matrix(C11, C12, C21, C22)

def pad_matrix(matrix, new_size):
    """
    Pad a matrix with zeros to make it a square matrix of size new_size x new_size.
    """
    old_size = len(matrix)
    padded_matrix = [[0] * new_size for _ in range(new_size)]

    for i in range(old_size):
        for j in range(old_size):
            padded_matrix[i][j] = matrix[i][j]

    return padded_matrix

def strassen_multiply(A, B):
    """
    Multiply two matrices A and B using Strassen's algorithm.
    Handles non-power-of-two matrix sizes by padding.
    """
    assert len(A) == len(A[0]) == len(B) == len(B[0]), "A and B must be square matrices"

    # Find the size of the matrices
    size = len(A)

    # If size is not a power of two, find the next power of two and pad the matrices
    new_size = 1
    while new_size < size:
        new_size *= 2

    if new_size != size:
        A = pad_matrix(A, new_size)
        B = pad_matrix(B, new_size)

    # Perform Strassen multiplication
    result = strassen(A, B)

    # Remove padding if necessary
    return [row[:size] for row in result[:size]]

# Example usage:
if __name__ == "__main__":
    A = [
        [1, 2, 0, 0],
        [3, 4, 0, 0],
        [0, 0, 1, 2],
        [0, 0, 3, 4]
    ]

    B = [
        [5, 6, 0, 0],
        [7, 8, 0, 0],
        [0, 0, 5, 6],
        [0, 0, 7, 8],
    ]

    result = strassen_multiply(A, B)
    print("Result of Strassen's Matrix Multiplication:")
    for row in result:
        print(row)
