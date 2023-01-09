package main

import (
	"log"
	"strings"
	"sync"
	"time"
)

const (
	INACTIVATED = "inactivated"
	ACTIVATED   = "activated"
	NOT_FOUND   = "status not found"
)

var statusEnum = [2]string{
	"inactivated",
	"activated",
}

// THIS IS THE WORST CODE IN THE WHOLE WORLD
func main() {
	PrintSomeInfos()
}

// printSomeInfos print some infos
func PrintSomeInfos() UserResponse {
	var mu sync.Mutex
	mu.Lock()

	timezone, _ := time.LoadLocation("America/New_York")

	var user UserResponse

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		user = UserResponse{lastName: getLastName("leonardo miranda"), age: getUserAge(time.Date(1996, 9, 02, 0, 0, 0, 0, timezone)), status: statusEnum[1]}
		wg.Done()
	}()

	wg.Wait()

	log.Println(user.age)
	log.Println(user.lastName)
	log.Println(user.status)

	mu.Unlock()

	return user
}

type User struct {
	name         string    // name
	birthdayDate time.Time // birthDate
	status       int       // status
}

type UserResponse struct {
	lastName string // last name
	age      int    // birthDate
	status   string // status
}

// this code run calculation of age and get the last name
func getInfos(u User) (int, string, string) {
	age := getUserAge(u.birthdayDate)

	lastName := getLastName(u.name)

	if u.status == 0 {
		return age, lastName, INACTIVATED
	}
	if u.status == 1 {
		return age, lastName, ACTIVATED
	}

	return age, lastName, NOT_FOUND
}

func getUserAge(birthdayDate time.Time) int {
	return int(time.Now().Sub(birthdayDate).Hours() / 24 / 365)
}

func getLastName(name string) string {
	splittedName := strings.Split(name, " ")
	return splittedName[len(splittedName)-1]
}
