package main

import (
	"fmt"
	"strings"
)

const (
	TESOURA = "tesoura"
	PEDRA   = "pedra"
	PAPEL   = "papel"
)

type game struct {
	name   string
	choice string
}

func main() {
	game1 := game{name: "Gabriel", choice: "tesoura"}
	game2 := game{name: "Alessandro", choice: "papel"}

	fmt.Println(play(game2, game1))
}

func play(game1, game2 game) string {

	var choices = make(map[string]int)
	choices[TESOURA] = 0
	choices[PEDRA] = 1
	choices[PAPEL] = 2

	if isDraw(strings.ToLower(game1.choice), strings.ToLower(game2.choice)) {
		return "empatou"
	}

	if strings.ToLower(game1.choice) == PAPEL && game2.choice == TESOURA {
		return game2.name
	}

	if game1.choice == TESOURA && game2.choice == PAPEL {
		return game1.name
	}

	if choices[game1.choice] > choices[game2.choice] {
		return game1.name
	} else {
		return game2.name
	}
}

func isDraw(choice1, choice2 string) bool {
	return choice1 == choice2
}

func comp(a, b int) {
	// 1 > 0
	// 1 > 2
	// 0 > 2
}
