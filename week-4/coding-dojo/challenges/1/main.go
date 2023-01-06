package main

import (
	"log"
	"strings"
	"sync"
	"time"
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

	a, ln, sn := 0, "", ""

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		a, ln, sn = getInfos(U{
			n:  "leonardo miranda",
			bd: time.Date(1996, 9, 02, 0, 0, 0, 0, timezone),
			s:  1,
		})
		wg.Done()
	}()

	wg.Wait()

	log.Println(a)
	log.Println(ln)
	log.Println(sn)

	mu.Unlock()
}

type U struct {
	n  string    // name
	bd time.Time // birthDate
	s  int       // status
}

// this code run calculation of age and get the last name
func getInfos(u U) (int, string, string) {
	a := int(time.Now().Sub(u.bd).Hours() / 24 / 365)

	fn := strings.Split(u.n, " ")

	if u.s == 0 {
		return a, fn[1], "inactivated"
	}
	if u.s != 0 && u.s == 1 {
		return a, fn[1], "activated"
	}
	if u.s != 0 && u.s != 1 {
		return a, fn[1], "status not found"
	}

	return a, fn[1], "status not found"
}
