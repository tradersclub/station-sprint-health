# Challenge 3

Dado uma lista de usuários com nome e ano de nascimento, crie uma função que receba a list usuários e calcule a idade deles e outra função que receba a lista de usuários com a idade ja calculada e remova os que tiverem menos que 18 anos. Crie testes unitários e de performance!


## Tipagem

Golang
```golang
    type User struct {
        name      string
        birthYear int
        age       int
    }
```

Typescript
```ts
    type User = {
        name:string;
        birthYear:number;
        age:number;
    }
```

## Exemplo

Input:
```js
// lista de usuários
[
    {
        name: "leo",
        birthYear: 1996
    },
    {
        name: "robertin",
        birthYear: 2000
    },
    {
        name: "theo",
        birthYear: 2018
    },
]
```

Output função calcular idade:
```js
// lista de usuários
[
    {
        name: "leo",
        birthYear: 1996,
        age: 27
    },
    {
        name: "robertin",
        birthYear: 2000,
        age: 23,
    },
    {
        name: "theo",
        birthYear: 2018,
        age: 5
    },
]
```

Output função remover menores:
```js
// lista de usuários
[
    {
        name: "leo",
        birthYear: 27
    },
    {
        name: "robertin",
        birthYear: 23
    },
]
```