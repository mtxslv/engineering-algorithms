Existem as seguintes quantidades de eleitores registrados:
	100 000 na zona urbana
	200 000 na zona suburbana
	 50 000 na zona rural
	 
É preciso de pelo menos a metade dos eleitores em cada zona. Logo:
	 50 000 na zona urbana
	100 000 na zona suburbana
	 25 000 na zona rural
	 

Temas a serem abordados nas campanhas:
	Apocalipse Zumbi
	Armar Tubarões com Lazers
	Avenidas para Carros Voadores
	
De acordo com a equipe de campanha, é possível estimar a quantidade de votos ganhados/vencidos por cada $1000 investido em tema de campanha. Os detalhes de veem abaixo:


|___________________|____________ZONA____________|
| TEMA DE CAMPANHA  | URBANA | SUBURBANA | RURAL |
|-------------------|--------|-----------|-------|
| Apocalipse Zumbi  |  -2    |     5     |   3   |
| Tubarões c/ lazers|   8    |     2     |  -5   |
| Avenida p/ car.voa|   0    |     0     |  10   |
| Voto de golfinho  |  10    |     0     |  -2   |
|-------------------|--------|-----------|-------|

Dessa forma, se você empregar a seguinte estratégia de campanha:
       $20 000 pro apocalipse zumbi
	    $0 pros lazers de tubarão
	$4 000  pra avenidas de carros voadores
	$9 000 pra votos de golfinhos

você terá a seguinte configuração de votos:
                     _           _          
_                   _|-2   5    3|   _                    _       
| 20000 0 4000 9000 || 8   2   -5| = | 50000 100000 82000 |       
                     | 0   0   10|          
                     |10   0   -2|            
 
Isto é, haverão

        50 000 votos urbanos ganhos
       100 000 votos suburbanos ganhos	
        82 000 votos rurais ganhos

Ou seja, o investimento de $33 000, distribuído nas áreas conforme mencionado acima, garantiu a quantidade de votos necessária (mínimo na metade do total por zona habitacional) 
        
A questão mais natural é: qual seria a quantidade mínima de dinheiro investida em cada tema de campanha capaz de garantir a quantidade mínima de votos por zona habitacional?

Para fazer isso é necessário "matematizar" nosso problema. Ou seja, gerar um modelo, ou uma descrição formal do que precisa ser realizado.

Em primeiro lugar, precisamos listar os inputs. Em outras palavras, "o que podemos manipular/variar" (nomeadas variáveis de decisão). Se nosso problema fosse uma caixa preta, o que estaríamos injetando?

  $$$   +----------+
------->| PROBLEMA |-------> votos
        +----------+
        
controlamos o dinheiro que vai pra cada tema de campanha. Essas serão as variáveis que definiremos:
	X1 são os milhares de $ pra campanha do Apocalipse Zumbi
	X2 são os milhares de $ pra campanha dos Tubarões com Lazers
	X3 são milhares de $ pra campanha das Rodovias de Carro Voador
	X4 são os milhares pra campanha dos Votos dos Golfinhos
	
A partir disso definimos quais nossas restrições? Bem, precisamos de uma quantidade mínima de votos em todas as 3 zonas. Mas como expressar isso em função de X1,X2,X3,X4?

	Votos Zona Urbana >=  50 000
	Votos Suburbanos  >= 100 000
	Votos Rurais      >=  25 000
	
Mas
	Votos Zona Urbana = -2X1 + 8X2 + 0X3 + 10X4
	 Votos Suburbanos =  5X1 + 2X2 + 0X3 + 0X4
	     Votos Rurais =  3X1 + -5X2 + 10X3 + -2X4
Portanto
         -2X1 + 8X2 +  0X3 + 10X4 >=  50 000	
          5X1 + 2X2 +  0X3 +  0X4 >= 100 000 
          3X1 - 5X2 + 10X3 -  2X4 >=  25 000

O que implica que cada conjunto de valores para X1,X2,X3,X4 que satisfaz as três inequações acima resolve o problema (quantidade de votos nesse caso). 

Voltando pro objetivo original, queremos minimizar a quantia investida. A expressão analítica para esse problema é a soma aplicada a cada tema. Logo:

         X1 + X2 + X3 + X4
         
Detalhe: uma vez que estamos investindo dinheiro, precisamos garantir que cada termo é maior que 0. Isso leva a novas restrições:

	X1 >= 0
	X2 >= 0
	X3 >= 0
	X4 >= 0 

Temos agora tudo o que é preciso pra definir formalmente nosso problema.

min{ X1 + X2 + X3 + X4 } |
         -2X1 + 8X2 +  0X3 + 10X4 >=  50 000	
          5X1 + 2X2 +  0X3 +  0X4 >= 100 000 
          3X1 - 5X2 + 10X3 -  2X4 >=  25 000
           X1                     >=       0
                 X2               >=       0
                        X3        >=       0
                               X4 >=       0 














        
        
        
        