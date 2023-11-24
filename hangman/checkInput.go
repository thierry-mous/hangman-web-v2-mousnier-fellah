package hangman

import (
	"strings"
)

func (g *Game) CheckInput(value string) string {
	value = strings.ToLower(value)
	if len(value) != 1 {
		if g.Word == value {
			for index := range g.Word {
				g.FoundLetters[index] = string(g.Word[index])
			}
			return "vous avez trouver le mot"
		} else {
			g.TurnsLeft -= 2
			return "vous avez perdu deux vies"
		}
	} else {

		for _, letter := range g.UsedLetters {
			if letter == value {
				return "vous avez deja indiquer cette lettre"
			}
		}
		g.UsedLetters = append(g.UsedLetters, value)
		IsFind := false
		for i, v := range g.Word {
			if value == string(v) {
				IsFind = true
				g.FoundLetters[i] = string(v)
			}
		}
		if !IsFind {
			g.TurnsLeft -= 1
			return "vous avez perdu une vie"
		}
		return "vous avez trouver la lettre"
	}
}