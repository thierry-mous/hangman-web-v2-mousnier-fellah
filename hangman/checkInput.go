package hangman

func (g *HangmanGame) CheckInput(value string) string {
	if len(value) != 1 {
		if g.word == value {
			for index := range g.word {
				g.foundLetter[index] = string(g.word[index])
			}
			return "vous avez trouvez le mot"
		} else {
			g.pv -= 1
			return " vous avez perdu 1 pv"
		}
	} else {
		IsFind := false
		for i, v := range g.word {
			if value == string(v) {
				IsFind = true
				g.hiddenWord[i] = string(v)
			}
		}
		if !(IsFind) {
			g.pv -= 1
			return "vous avez perdu 1 pv"
		}
		return "vous avez trouvez le mot"
	}
}
