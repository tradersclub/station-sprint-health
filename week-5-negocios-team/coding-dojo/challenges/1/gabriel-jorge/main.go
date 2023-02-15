package main

type User struct {
	id          string
	isActivated bool
}

const inputTeste := [
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

func main() {

}

func compareLists(userList, userMasterList []User) []User {
	mapUserList := make(map[string]bool)
	data := []User{}

	for _, user := range userList {
		if user.isActivated {
			mapUserList[user.id] = user.isActivated
		}
	}

	for _, user := range userMasterList {
		if !mapUserList[user.id] {
			data = append(data, user)
		}
	}

	return data
}
