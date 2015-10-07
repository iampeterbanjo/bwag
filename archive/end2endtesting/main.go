package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func HelloWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(res, "Hello World")
}

func App() http.Handler {
	n := negroni.New()

	m := func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
		fmt.Fprint(res, "Before...")
		next(res, req)
		fmt.Fprint(res, "...After")
	}
	n.Use(negroni.HandlerFunc(m))

	r := httprouter.New()

	r.GET("/", HelloWorld)
	n.UseHandler(r)

	return n
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("ListenAndServe on ", port)
	http.ListenAndServe(":" + port, App())
}