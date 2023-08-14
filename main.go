package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func h1(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "Casablanca", Director: "Michael Curtiz"},
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
		},
	}

	err := tmpl.Execute(w, films)
	if err != nil {
		log.Fatal(err)
	}
}

func h2(w http.ResponseWriter, r *http.Request) {
	// this simulates a delay on the server
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	htmlStr := fmt.Sprintf("<li>%s: %s</li>", title, director)
	templ, err := template.New("t").Parse(htmlStr)
	if err != nil {
		log.Fatal(err)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("listening on port 8000")

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
