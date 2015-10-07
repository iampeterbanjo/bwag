package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
)

func Test_GetHelloWorld(t *testing.T) {
	domain, port := os.Getenv("domain"), os.Getenv("port")
	if domain == "" {
		domain = "http://localhost"
	}

	if port == "" {
		port = "3000"
	}

	req, err := http.NewRequest("GET", domain + port, nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	HelloWorld(res, req, nil)

	exp := "Hello World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}

func Test_PostHelloWorld(t *testing.T) {
	domain, port := os.Getenv("domain"), os.Getenv("port")
	if domain == "" {
		domain = "http://localhost"
	}

	if port == "" {
		port = "3000"
	}

	req, err := http.NewRequest("POST", domain + port, nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	GoodbyeWorld(res, req, nil)

	exp := "Goodbye World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}