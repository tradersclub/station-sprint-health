
import { shuffleList } from "./utils";

const listPerson: string[] = ["Roberval", "Cleiton",  "Patricia", "Bernardo"];
const qtyPersonByTeam = 2;
const qtyTeams = listPerson.length / qtyPersonByTeam;

function createTeam(list: string[], sizeTeam: number){
    const newList = [];

    for(let i = 0; i < sizeTeam; i++){
        newList.push(list[i])
    }
    return newList
}

function createTeams () {
    const randomTeam = shuffleList(listPerson);

    const newTeams: string[][] = []
    let list = randomTeam;

    for (let i = 0; i < qtyTeams; i++){
        const newTeam: string[] = createTeam(list, qtyPersonByTeam);
        newTeams.push(newTeam)
        list = list.filter(item => !newTeam.includes(item))
    }

    return newTeams
};

console.log(createTeams())
