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
		negroni.HandlerFunc(LetMeGoogleThatForYou),
		negroni.HandlerFunc(MyMiddleware),
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

func LetMeGoogleThatForYou(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	question := r.URL.Query().Get("question")
	redirectUrl := "http://lmgtfy.com/?q="

	if question != "" {
		http.Redirect(rw, r, redirectUrl + question, http.StatusTemporaryRedirect)
	} else {
		next(rw, r)
	}
}