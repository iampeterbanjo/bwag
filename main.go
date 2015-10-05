package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type FamilyMember struct {
	Name string
	Age int
	Parents []string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/books", ShowBooks)
	http.HandleFunc("/family", ShowFamily)
	http.ListenAndServe(":" + port, nil)
}

func ShowFamily(w http.ResponseWriter, r *http.Request) {
	family := FamilyMember{"Mystery Banjo", 24, []string{"Mum","Dad"}}
	js, err := json.Marshal(family)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	js, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
