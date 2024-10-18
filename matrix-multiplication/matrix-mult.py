def regular_matrix_mult(matrix_a, matrix_b, matrix_c, n):
    for i in range(0, n):
        for j in range(0, n):
            for k in range(0,n):
                matrix_c[i][j] = matrix_c[i][j] + matrix_a[i][k]*matrix_b[k][j] 