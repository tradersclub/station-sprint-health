package main

import (
	"fmt"
	"time"
)

type User struct {
	name      string
	birthYear int
	age       int
}

var usersTest = []User{
	{
		name:      "leo",
		birthYear: 1996,
	},
	{
		name:      "robertin",
		birthYear: 2000,
	},
	{
		name:      "theo",
		birthYear: 2018,
	},
	{
		name:      "neto",
		birthYear: 1991,
	},
	{
		name:      "matheus",
		birthYear: 1994,
	},
}

func main() {
	usrs := CalculatePersonAge(usersTest)
	usrs18OrMore := Get18OrMore(usrs)

	fmt.Println(usrs)
	fmt.Println(usrs18OrMore)
}

func CalculatePersonAge(users []User) []User {
	currentYear := time.Now().Year()

	for i, usr := range users {
		users[i].age = currentYear - usr.birthYear
	}

	return users
}

func Get18OrMore(users []User) []User {
	users18 := make([]User, 0, len(users))

	for _, usr := range users {
		if usr.age >= 18 {
			users18 = append(users18, usr)
		}
	}

	return users18
}
