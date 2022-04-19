package main

import "fmt"

type jogo struct {
	name   string
	choice string
}

func main() {
	jogo1 := jogo{name: "Gabriel", choice: "tesoura"}
	jogo2 := jogo{name: "Alessandro", choice: "papel"}

	fmt.Println(play(jogo2, jogo1))
}

func play(jogo1, jogo2 jogo) string {
	var choices = make(map[string]int)
	choices["tesoura"] = 0
	choices["pedra"] = 1
	choices["papel"] = 2

	if jogo1.choice == jogo2.choice {
		return "empatou"
	}

	if jogo1.choice == "papel" && jogo2.choice == "tesoura" {
		return jogo2.name
	}

	if jogo1.choice == "tesoura" && jogo2.choice == "papel" {
		return jogo1.name
	}

	if choices[jogo1.choice] > choices[jogo2.choice] {
		return jogo1.name
	} else {
		return jogo2.name
	}
}
