package main

import (
	"fmt"
	"math/rand"
	"time"
)

var stationDevelopers = []string{
	"Wagner",
	"Gabriel",
	"Joabson",
	"Guilherme",
	"Luciano",
	"Alessandro",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	developers := stationDevelopers
	driver, coDriver := "", ""

	totalTeams := len(developers) / 2

	for i := 0; i < totalTeams; i++ {

		driver, developers = pickRandomDeveloper(developers)

		coDriver, developers = pickRandomDeveloper(developers)

		printTeam(driver, coDriver)

	}

}

func pickRandomDeveloper(developers []string) (driver string, remainingDevelopers []string) {

	totalDevelopers := len(developers)
	driverIndex := rand.Intn(totalDevelopers)
	driver = developers[driverIndex]

	remainingDevelopers = removeIndexFromSlice(developers, driverIndex)

	return
}

func printTeam(driver, coDriver string) {
	fmt.Println("\nTeam:")
	fmt.Println("Piloto:", driver)
	fmt.Println("Co-piloto:", coDriver)
}

func removeIndexFromSlice(slice []string, index int) []string {
	lastIndex := len(slice) - 1

	slice[index] = slice[lastIndex]
	slice = slice[:lastIndex]

	return slice
}
