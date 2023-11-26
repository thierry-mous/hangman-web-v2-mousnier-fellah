package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Game struct {
	State        bool
	Letters      []string // nul
	FoundLetters []string
	UsedLetters  []string
	Word         string
	TurnsLeft    int
	InGame       bool
}

var Player Game

func New(turns int, word string) (*Game, error) {
	if len(word) < 2 {
		return nil, fmt.Errorf("Le mot '%s' doit faire minimum 2 charactères . got=%v", word, len(word))
	}

	letters := strings.Split(strings.ToLower(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        true,
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
		Word:         strings.ToLower(word),
		InGame:       true,
	}

	return g, nil
}

func HasWon(hiddenWord []string, word string) bool {
	if len(hiddenWord) == len(word) {
		for i := range hiddenWord {
			if string(hiddenWord[i]) != string(word[i]) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

/*func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}

func (g *Game) LoseTurn(guess string) {
	g.TurnsLeft--
	g.UsedLetters = append(g.UsedLetters, guess)
}

func LetterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
} */

var words = make([]string, 0, 50)

func PrepareFileName(level string) string {
	switch level {
	case "1":
		return "level1.txt"
	case "2":
		return "level2.txt"
	case "3":
		return "level3.txt"
	case "4":
		return "level4.txt"
	case "5":
		return "level5.txt"
	case "6":
		return "level6.txt"
	case "7":
		return "levelbonus.txt"
	default:
		return "level1.txt"
	}
}

func Load(filename string) error { //charge le .txt avec les noms de spider
	f, err := os.Open("./hangman/data/" + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func PickWord() string { //prend un nom de spider aléatoire
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}

/*var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("Quelle est votre lettre? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		guess = strings.TrimSpace(guess)

		if len(guess) != 1 {
			fmt.Println("Format invalide.", guess, len(guess))
			continue
		}
		valid = true
	}
	return
} */

func Start(level string) {
	err := Load(PrepareFileName(level))
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	g, err := New(5, PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}
	Player = *g
	fmt.Println("game init word : ", g.Word)
}
