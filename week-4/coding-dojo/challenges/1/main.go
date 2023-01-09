package main

import (
	"log"
	"strings"
	"time"
)

// criar teste
// limpar codigo
// redesenhar o style

// THIS IS THE WORST CODE IN THE WHOLE WORLD
func main() {
	printSomeInfos()
}

// printSomeInfos print some infos
func printSomeInfos() {

	timezone, _ := time.LoadLocation("America/New_York")
	a, ln, sn := getInfos(User{
		name:     "leonardo miranda",
		birthday: time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
		status:   1,
	})

	log.Println(a)
	log.Println(ln)
	log.Println(sn)

}

type User struct {
	name     string
	birthday time.Time
	status   int
}

const hoursInOneDay = 24
const daysInOneYear = 365

func (user User) Age() int {
	return int(time.Now().Sub(user.birthday).Hours() / hoursInOneDay / daysInOneYear)
}

func (user User) Lastname() string {
	splittedName := strings.Split(user.name, " ")
	return splittedName[len(splittedName)-1]
}

func (user User) StatusToString() string {
	switch user.status {
	case 0:
		return "inactivated"
	case 1:
		return "activated"
	default:
		return "status not found"
	}
}

// this code run calculation of age and get the last name
func getInfos(user User) (int, string, string) {
	return user.Age(), user.Lastname(), user.StatusToString()
}
