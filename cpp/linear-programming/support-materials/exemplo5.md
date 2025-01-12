# Original Problem:

$ T^{(0)} = 
\begin{bmatrix}
1 & -2 & -3 & 0 & 0 & 0  &  0 \\
0 &  1 &  0 & 1 & 0 & 0  &  3 \\
0 &  0 &  1 & 0 & 1 & 0  &  4 \\
0 &  1 &  3 & 0 & 0 & 1  & 12 \\
\end{bmatrix}
\implies q^{(0)} = \begin{bmatrix}
\infty \\
4 \\
4
\end{bmatrix}
$

# Algoritmo rodando..

Escolhendo como pivot o elemento $T_{3,3}^{(0)}$, temos as operações:

$L_1 \leftarrow L_1 + 3L_3$

$L_4 \leftarrow L_4 - 3L_3$


$ T^{(1)} = 
\begin{bmatrix}
1 & -2 & 0 & 0  &  3 & 0  &  12 \\
0 &  1 &  0 & 1 &  0 & 0  &   3 \\
0 &  0 &  1 & 0 &  1 & 0  &   4 \\
0 &  1 &  0 & 0 & -3 & 1  &   0 \\
\end{bmatrix}
\implies q^{(1)} = \begin{bmatrix}
3 \\
\infty \\
0
\end{bmatrix}
$

Como a linha é aquela em que temos o menor quociente positivo, iremos escolher a linha 2. Assim:

$L_1 \leftarrow L_1 + 2L_2$

$L_4 \leftarrow L_4 - L_1$

$ T^{(2)} = 
\begin{bmatrix}
1 &  0 & 0 &  2 &  3 & 0  &  18 \\
0 &  1 & 0 &  1 &  0 & 0  &   3 \\
0 &  0 & 1 &  0 &  1 & 0  &   4 \\
0 &  0 & 0 & -1 & -3 & 1  &  -3 \\
\end{bmatrix}
$


Observe que os coeficiente da linha 1 estão todos positivos. Paramos por aqui. Assim:

$
\begin{bmatrix}
x_1 \\
x_2
\end{bmatrix}
= \begin{bmatrix}
3 \\
4 
\end{bmatrix};
\begin{bmatrix}
s_1 \\
s_2 \\
s_3 
\end{bmatrix}
= \begin{bmatrix}
0 \\
0 \\
-3
\end{bmatrix};
z = \begin{bmatrix} 2 & 3\end{bmatrix}
\begin{bmatrix} 3 \\ 4 \end{bmatrix} = 2\cdot3 + 3\cdot 4 = 6 + 12 = 18
$

Mas não faz sentido termos uma variável de folga negativa! Creio que esse exemplo tá errado.