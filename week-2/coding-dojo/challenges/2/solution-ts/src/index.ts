import { createUsers, makeMapFromUsersArray } from './utils';

const users = createUsers(10**6);
const mapUsers = makeMapFromUsersArray(users)

console.log(`map users size: ${mapUsers.size}`)

console.time("delete")
mapUsers.delete(1)
console.timeEnd("delete")

console.time(`set`)
mapUsers.set(1,{ id:1, name:`name_add` })
console.timeEnd("set")

console.time("get")
mapUsers.get(500)
console.timeEnd("get")

