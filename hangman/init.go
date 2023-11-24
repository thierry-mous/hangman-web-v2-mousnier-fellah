package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type HangmanGame struct {
	isGame      bool
	word        string
	hiddenWord  []string //mot cache
	listLetter  []string // lettre deja mise
	pv          int
	isWin       bool
	foundLetter []string
}

var Player HangmanGame

func contains(arr []int, elem int) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}

func (g *HangmanGame) Init(fileName string) {
	path := "./hangman/data/" + fileName
	fichier, err := os.Open(path)
	if err != nil {
		fmt.Println("Oupsss une erreur de lecture du fichier...")
		os.Exit(1)
	}
	defer fichier.Close()
	var mots []string
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Oupsss une erreur de scan du fichier...")
		os.Exit(1)
	}
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	indexAleatoire := rng.Intn(len(mots))
	g.word = mots[indexAleatoire]

	for i := 0; i < len(g.word); i++ {
		g.hiddenWord = append(g.hiddenWord, "_")
	}

	indices := make([]int, 0)

	for len(indices) < 2 {
		i := rand.Intn(len(g.word))
		// Vérifier si l'indice n'a pas déjà été choisi et si la lettre correspondante ne se répète pas dans le mot
		if !contains(indices, i) {
			indices = append(indices, i)
		}
	}
	g.Initword(indices)
	g.isGame = true
	g.pv = 5
	g.isWin = false
}

func (g *HangmanGame) Initword(listeIndice []int) {
	for _, indice := range listeIndice {
		for index, value := range g.word {
			if string(value) == string(g.word[indice]) {
				g.hiddenWord[index] = string(value)
			}
		}
	}
}
