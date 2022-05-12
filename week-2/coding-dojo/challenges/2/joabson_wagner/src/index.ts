type User = {
    id: number;
    name: string;
}
const lengthList = 1000000;
const createUser = (_: number, index: number) => {
    const user: User = {
        id: index,
        name: 'user_' + index
    }
    return user
}
const searchUser = (list: User[], id: number): User | undefined => {
    return list.find(user => user.id === id)
}
const deleteUser = (list: User[], id: number): User[] => {
    return list.filter(user => user.id !== id)
}
const generateList = (length: number): User[] => {
    const list = Array.from({ length }, createUser)
    return list
}


const main = () => {
    const listUsers = generateList(lengthList)
    console.log(`:foguete: ~ listUsers`, listUsers.length)
    console.time('search')

    searchUser(listUsers, lengthList)
    console.timeEnd('search')
    console.time('deleteUser')
    const newList = deleteUser(listUsers, 1)
    // console.log(`:foguete: ~ delete`, newList.length)
    console.timeEnd('deleteUser')
    console.time('createUser')
    const userTest = createUser(10, listUsers.length + 1);
    // console.log(`:foguete: ~ userTest`, userTest)
    console.timeEnd('createUser')
}
main()