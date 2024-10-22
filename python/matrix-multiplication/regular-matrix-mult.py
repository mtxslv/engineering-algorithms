import time
import random

how_many_trials = 10
tempos_4x4 = [0 for _ in range(0,how_many_trials)]
it_4x4 = 0

def regular_matrix_mult(matrix_a, matrix_b, matrix_c, n):
    global it_4x4
    start_time = time.time_ns()
    for i in range(0, n):
        for j in range(0, n):
            for k in range(0,n):
                matrix_c[i][j] = matrix_c[i][j] + matrix_a[i][k]*matrix_b[k][j] 
    end_time = time.time_ns()
    total_nanosec_time = end_time - start_time
    tempos_4x4[it_4x4] = total_nanosec_time
    
    it_4x4 = it_4x4 + 1


def generate_matrix(n, fill=None):
    param = 10
    matrix = []

    for i in range(0,n):
        row = [
            random.randint(-param, param) if fill is None else fill for i in range(0,n)
        ]
        matrix.append(row)
    return matrix

if __name__ == '__main__':
    for trial in range(0, how_many_trials):
        matrix_A = generate_matrix(4)
        matrix_B = generate_matrix(4)
        matrix_C = generate_matrix(4,0)
        regular_matrix_mult(
            matrix_A,
            matrix_B,
            matrix_C,
            4
        )
    print(tempos_4x4)
    # a_matrix = generate_matrix(4)
    # for row in a_matrix:
    #     print(row)
    # matrix_a = [
    #     [1, 2, 0, 0],
    #     [3, 4, 0, 0],
    #     [0, 0, 1, 2],
    #     [0, 0, 3, 4]
    # ]
    # matrix_b = [
    #     [5, 6, 0, 0],
    #     [7, 8, 0, 0],
    #     [0, 0, 5, 6],
    #     [0, 0, 7, 8]
    # ]    
    # matrix_c = [
    #     [0, 0, 0, 0],
    #     [0, 0, 0, 0],
    #     [0, 0, 0, 0],
    #     [0, 0, 0, 0]
    # ] 
    # expected = [
    #     [19, 22,  0,  0],
    #     [43, 50,  0,  0],
    #     [ 0,  0, 19, 22],
    #     [ 0,  0, 43, 50]
    # ]
    # regular_matrix_mult(matrix_a, matrix_b, matrix_c, 4)
    
    # for row in matrix_c:
    #     print(row)    