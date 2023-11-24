package hangman

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Route() { //lire tout nos templates html

	tmpl, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home", nil)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "level", nil)
	})

	http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "result", nil)
	})

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "menu", nil)

	})

	http.HandleFunc("/choice", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "level", nil)
	})

	//lie le css
	rootDoc, _ := os.Getwd()
	fmt.Println("Serveur écoutant sur le port 8085")
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8085", nil)

}
