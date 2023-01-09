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

// THIS IS THE WORST CODE IN THE WHOLE WORLD
func main() {
	printSomeInfos()
}

// printSomeInfos print some infos
func printSomeInfos() {
	var mu sync.Mutex
	mu.Lock()

	timezone, _ := time.LoadLocation("America/New_York")

	age, lastName, status := 0, "", ""

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		age, lastName, status = getInfos(User{
			name:         "leonardo miranda",
			birthdayDate: time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
			status:       1,
		})
		wg.Done()
	}()

	wg.Wait()

	log.Println(age)
	log.Println(lastName)
	log.Println(status)

	mu.Unlock()
}

type User struct {
	name         string    // name
	birthdayDate time.Time // birthDate
	status       int       // status
}

// this code run calculation of age and get the last name
func getInfos(u User) (int, string, string) {
	age := getUserAge(u.birthdayDate)

	lastName := getLastName(u.name)

	if u.status == 0 {
		return age, lastName, INACTIVATED
	}
	if u.status != 0 && u.status == 1 {
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
