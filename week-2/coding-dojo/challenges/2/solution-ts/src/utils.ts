import {MapUsers, User} from './types';

export function createUsers(qty:number):User[] {
    let users:User[] = []
    
    for (let index = 0; index < qty; index++) {
        users[index] = {
            id: index,
            name:`name${index}`
        };
    }
    
    return users;
}

export function makeMapFromUsersArray(users:User[]):MapUsers {
    const mapUsers = new Map();

    users.forEach((users) => {
        mapUsers.set(users.id, users);
    })

    return mapUsers;
}