package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Middleware stack
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.HandlerFunc(MyMiddleware),
		negroni.HandlerFunc(LetMeGoogleThatForYou),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("public")),
	)

	n.Run(":" + port)
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("Logging on the way there")

	if r.URL.Query().Get("password") == "secret123" {
		next(rw, r)
	} else {
		http.Error(rw, "Not Authorized", 401)
	}

	log.Println("Logging on the way back")
}

func LetMeGoogleThatForYou(rw http.ResponseWriter, r *http.Request, http.HandlerFunc) {
	if r.URL.Query().Get("question") != "" {
		r.redirect("www.google.com")
	} else {
		next(rw, r)
	}
}