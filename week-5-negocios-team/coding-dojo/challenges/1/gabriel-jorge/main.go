package main

import "fmt"

type User struct {
	id          string
	isActivated bool
}

var Users = []User{
	{
		id:          "1",
		isActivated: true,
	},
	{
		id:          "2",
		isActivated: false,
	},
	{
		id:          "3",
		isActivated: true,
	},
}

var UsersMaster = []User{
	{
		id:          "3",
		isActivated: true,
	},
}

func main() {
	res := CompareLists(Users, UsersMaster)

	fmt.Println(res)
}

func CompareLists(userList, userMasterList []User) []User {
	mapUserList := make(map[string]bool)
	data := []User{}

	for _, user := range userMasterList {
		if user.isActivated {
			mapUserList[user.id] = user.isActivated
		}
	}

	for _, user := range userList {
		if !mapUserList[user.id] && user.isActivated {
			data = append(data, user)
		}
	}

	return data
}
