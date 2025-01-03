```shell
$ make
$ ./simplexsolver example2.txt 2 4
```


[Example 1](https://math.libretexts.org/Bookshelves/Applied_Mathematics/Applied_Finite_Mathematics_(Sekhon_and_Bloom)/04%3A_Linear_Programming_The_Simplex_Method/4.02%3A_Maximization_By_The_Simplex_Method)

[Example 2](https://personal.utdallas.edu/~scniu/OPRE-6201/documents/LP06-Simplex-Tableau.pdf)

[Example 3 comes from page 107](http://www.maths.lse.ac.uk/Personal/stengel/HillierLieberman9thEdition.pdf)

[Example 4 comes from page 29](./support-materials/Disciplina_Otimizao_de_Sistemas_-1a_parte_v9_corr3.pdf)

# EXPECTED ANSWERS

[example1.txt](./example1.txt) must lead a maximized objective of 400. Since the code was based on this example, the answer is correct.

[example2.txt](./example2.txt) leads to a value of 9. The code was tested against this example, so the answer is also correct.

[example3.txt](./example3.txt) gives a wrong maximized value. I don't know what might be wrong

[example4.txt](./example4.txt) should return a maximized 82.5, but my solution stops a step earlier (72)  

[example5.txt](./example5.txt) must give a maximized objective value of 15. The exact form of the answer may depend, since we have a degenerated solution.
