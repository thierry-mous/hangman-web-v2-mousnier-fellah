package hangman

type HangmanGame struct {
	isGame bool
	word string
	hiddenWord []string
	listLetter []string
	pv int
	isWin bool
}

