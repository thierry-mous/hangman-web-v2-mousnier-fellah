package main

import (
	"fmt"
	"hangmanweb/hangman"
	"html/template"
	"net/http"
	"os"
)

type PageInit struct {
	Username string
	lvl      string
}

type Game struct {
	State        string
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnsLeft    int
}

var logs PageInit
var MesUser string

func main() {
	temp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(1)
		temp.ExecuteTemplate(w, "home", nil)
	})

	http.HandleFunc("/choice", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(2)
		temp.ExecuteTemplate(w, "menu", nil)

	})

	http.HandleFunc("/choice/treatment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(3)
		logs = PageInit{r.FormValue("pseudo"), r.FormValue("level")}
		fmt.Println(logs)
		hangman.Start(logs.lvl)
		http.Redirect(w, r, "/game", 301)
	})
	type PageGame struct {
		Hiddenword []string
		Listletter []string
		Leftpv     int
		MesUser    string
		IsWin      bool
	}

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(4)
		data := PageGame{hangman.Player.FoundLetters, hangman.Player.UsedLetters, hangman.Player.TurnsLeft, MesUser, false}
		if hangman.HasWon(hangman.Player.FoundLetters, hangman.Player.Word) || hangman.Player.TurnsLeft <= 0 {
			hangman.Player.InGame = false
			http.Redirect(w, r, "/end", 301)
		}
		temp.ExecuteTemplate(w, "level", data)
	})

	http.HandleFunc("/game/treatment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(5)
		value := r.FormValue("value")
		fmt.Println("valeur niveau : ", value)
		MesUser = hangman.Player.CheckInput(value)
		data := PageGame{hangman.Player.FoundLetters, hangman.Player.UsedLetters, hangman.Player.TurnsLeft, MesUser, false}
		if hangman.HasWon(hangman.Player.FoundLetters, hangman.Player.Word) || hangman.Player.TurnsLeft <= 0 {
			hangman.Player.InGame = false
			http.Redirect(w, r, "/end", 301)
		} 
		temp.ExecuteTemplate(w, "level", data)
	})

	http.HandleFunc("/end", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(6)
		data := PageGame{hangman.Player.FoundLetters, hangman.Player.UsedLetters, hangman.Player.TurnsLeft, MesUser, true}
		temp.ExecuteTemplate(w, "end", data)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Init serv
	http.ListenAndServe("localhost:8080", nil)
}
