package main_rendering

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"log"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("listening on " + port)
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":" + port, nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
