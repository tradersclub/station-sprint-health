# Challenge 1

Dado uma lista de 500 usuários e 500 usuários com plano master. Crie um programa que descobre quem foram as pessoas que não tem plano master e estão ativas. Inclua testes unitários e explique a complexidade algoritmica da sua solução.


## Tipagem

Golang:
```go
type User struct {
    id string
    isActivated bool
}
```

Typescript:
```ts
type User = {
    id:string;
    isActivated:boolean;
}
```

## Exemplo

Input:
```js
// users
[
    {
        id:"1"
        isActivated: true,
    }
    {
        id:"2"
        isActivated: false,
    }
    {
        id:"3"
        isActivated: true,
    }
]

// usersMaster
[
    {
        id:"3"
        isActivated: true,
    }
]
```

Output:
```js
[
    {
        id:"1"
        isActivated: true,
    }
]
```