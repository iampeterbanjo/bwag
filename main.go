package main

import (
	"net/http"
	"os"
	"gopkg.in/unrolled/render.v1"
)

type MyController struct {
	AppController
	*render.Render
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}

// Action defines a standard function signature for us to use when
// creating controller actions. A controller action is basically just
// a method attached to a controller
type Action func(rw http.ResponseWriter, r *http.Request) error

// This is our base controller
type AppController struct{}

// The action function helps with error handling in a controller
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

func (c *MyController) Index(rw http.ResponseWriter, r *http.Request) error {
	c.JSON(rw, 200, map[string]string{"Hello": "JSON"})
	return nil
}

func main() {
	c := &MyController{Render: render.New(render.Options{})}
	http.ListenAndServe(":" + GetPort(), c.Action(c.Index))
}