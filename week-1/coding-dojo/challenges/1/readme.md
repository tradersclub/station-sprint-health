# Challenge 1
A partir de uma lista de usuários, com nome e data de aniversário, busque o NOME dos usuários com mais de 18 e que o nome é um palíndromo.

> Palíndromo: diz-se de ou frase ou palavra que se pode ler, indiferentemente, da esquerda para a direita ou vice-versa.

Exemplo 1:
```json
[
    {
        "name":"Juliana",
        "birthday":956026800000,
    },
    {
        "name":"Ana",
        "birthday":957927600000,
    }
]
```
Resultado: Ana

Exemplo 2:
```json
[
    {
        "name":"Juliana",
        "birthday":956026800000,
    },
    {
        "name":"Roberto",
        "birthday":957927600000,
    }
]
```
Resultado: nulo

Exemplo 3:
```json
[
    {
        "name":"Juliana",
        "birthday":956026800000,
    },
    {
        "name":"Natan",
        "birthday":1589079600000,
    }
]
```
Resultado: nulo

Gerador de timestamp: [https://www.epochconverter.com/](https://www.epochconverter.com/)