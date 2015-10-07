package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/julienschmidt/httprouter"
)

func HelloWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(res, "Hello World")
}

func GoodbyeWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprint(res, "Goodbye World")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := httprouter.New()
	r.GET("/", HelloWorld)
	r.POST("/", GoodbyeWorld)

	fmt.Println("ListenAndServe on " + port)
	http.ListenAndServe(":" + port, r)
}