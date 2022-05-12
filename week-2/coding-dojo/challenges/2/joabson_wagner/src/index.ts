type User = {
    id: number;
    name: string;
}

// função para gerar lista 1.000.000


// função para gerar 1 user
const createUser = (_index: number, index: number) => {
    const user: User = {
        id: index,
        name: 'user_' + index
    }

    return user
}

// função de busca
// const searchUser = (id: number): User | undefined => {
//     return list.find(user => user.id === id)
// }

// // função de delete
// const deleteUser = (id: number): User[] => {
//     return list.filter(user => user.id !== id)
// }

const generateList = (length: number) => {
    const list = Array.from({ length }, createUser)
    return list
}


const main = () => {
    const listUsers = generateList(10)
    console.log(`🚀 ~ listUsers`, listUsers)
}

main()