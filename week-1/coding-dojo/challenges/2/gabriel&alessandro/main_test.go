package main

import "testing"

func TestPlay(t *testing.T) {
	var (
		game1 = game{
			name:   "Gabriel",
			choice: "papel",
		}
		game2 = game{
			name:   "Alessandro",
			choice: "papel",
		}
		game3 = game{
			name:   "Luciano",
			choice: "tesoura",
		}
		game4 = game{
			name:   "Leo",
			choice: "pedra",
		}
	)

	out := play(game1, game2)
	if out != "empatou" {
		t.Error("resultado errado")
	}

	out = play(game1, game3)
	if out != "Luciano" {
		t.Error("resultado errado")
	}

	out = play(game1, game4)
	if out != "Gabriel" {
		t.Error("resultado errado")
	}

	out = play(game3, game4)
	if out != "Leo" {
		t.Error("resultado errado")
	}

}
