```shell
$ make
$ ./simplexsolver ./maximization-examples/example/example2.txt 2 4
```


[Example 1](https://math.libretexts.org/Bookshelves/Applied_Mathematics/Applied_Finite_Mathematics_(Sekhon_and_Bloom)/04%3A_Linear_Programming_The_Simplex_Method/4.02%3A_Maximization_By_The_Simplex_Method)

[Example 2](https://personal.utdallas.edu/~scniu/OPRE-6201/documents/LP06-Simplex-Tableau.pdf)

[Example 3 comes from page 107](http://www.maths.lse.ac.uk/Personal/stengel/HillierLieberman9thEdition.pdf)

[Example 4 comes from page 29](./support-materials/Disciplina_Otimizao_de_Sistemas_-1a_parte_v9_corr3.pdf)

# EXPECTED ANSWERS

## [example1.txt](./maximization-examples/example1.txt) 

must lead a maximized objective of 400. Since the code was based on this example, the answer is correct.

```shell
$ ./simplexsolver ./maximization-examples/example/example1.txt 2 2
Processed tableau containing 2 variable(s) and 2 constraint(s).
Maximized objective value: 400.00000
x1 = 4.00000
x2 = 8.00000
s1 = 0.00000
s2 = 0.00000
```

## [example2.txt](./maximization-examples/example2.txt)

leads to a value of 9. The code was tested against this example, so the answer is also correct.

```shell
$ ./simplexsolver ./maximization-examples/example/example2.txt 2 4
Processed tableau containing 2 variable(s) and 4 constraint(s).
Maximized objective value: 9.00000
x1 = 1.50000
x2 = 1.00000
s1 = 0.00000
s2 = 5.50000
s3 = 3.00000
s4 = 0.00000
```

## [example3.txt](./maximization-examples/example3.txt)

Answer correct.

```shell
$ ./simplexsolver ./maximization-examples/example/example3.txt 2 3
Processed tableau containing 2 variable(s) and 3 constraint(s).
Maximized objective value: 36.00000
x1 = 2.00000
x2 = 6.00000
s1 = 2.00000
s2 = 0.00000
s3 = 0.00000
```

## [example4.txt](./maximization-examples/example4.txt) 

Correct answer.

```shell
$ ./simplexsolver ./maximization-examples/example/example4.txt 2 4
Processed tableau containing 2 variable(s) and 4 constraint(s).
Maximized objective value: 82.50000
x1 = 2.50000
x2 = 3.50000
s1 = 0.50000
s2 = 0.50000
s3 = 0.00000
s4 = 0.00000
```

## [example5.txt](./maximization-examples/example5.txt) 

**WRONG ANSWER HERE! DO NOT CONSIDER THIS EXAMPLE**
must give a maximized objective value of 15. The exact values of the variables (slack and problem's) may depend, since we have a degenerated solution. Even so, it returns what is expected in [page 35](./support-materials/Disciplina_Otimizao_de_Sistemas_-1a_parte_v9_corr3.pdf).

```shell
$ ./simplexsolver ./maximization-examples/example/example5.txt 2 3
Processed tableau containing 2 variable(s) and 3 constraint(s).
Maximized objective value: 15.00000
x1 = 3.00000
x2 = 3.00000
s1 = 0.00000
s2 = 1.00000
s3 = 0.00000
```


## [example6.txt](./maximization-examples/example6.txt)

[página 47](http://www.maths.lse.ac.uk/Personal/stengel/HillierLieberman9thEdition.pdf)

## [example7.txt](./maximization-examples/example7.txt)

[página 4](https://egyankosh.ac.in/bitstream/123456789/18135/1/Unit-4.pdf)

## [example8.txt](./maximization-examples/example8.txt)

It comes from [this website](https://people.richland.edu/james/ictcm/2006/3dsimplex.html). The answer is correct:

```shell
$ ./simplexsolver ./maximization-examples/example/example8.txt 3 4
Processed tableau containing 3 variable(s) and 4 constraint(s).
Maximized objective value: 268.00000
x1 = 1.80000
x2 = 20.80000
x3 = 1.60000
s1 = 0.00000
s2 = 0.00000
s3 = 2.60000
s4 = 0.00000
```

## [example9.txt](./maximization-examples/example9.txt)

[examples comes from youtube example](https://youtu.be/9YKLXFqCy6E?si=dO8H4DqDfVpinkWu)

```shell
$ ./simplexsolver ./maximization-examples/example/example9.txt 2 2
Processed tableau containing 2 variable(s) and 2 constraint(s).
Maximized objective value: 32.00000
x1 = 2.00000
x2 = 3.00000
s1 = 0.00000
s2 = 0.00000
```