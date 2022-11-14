package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if host == "" {
		host = ""
	}
	if port == "" {
		port = "8080"
	}
	addr := host + ":" + port
	s := &Server{}
	http.Handle("/", s)

	log.Println("Listening on address", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type Server struct {
}

type Controller struct {
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &Controller{}
	switch r.URL.Path {
	case "/":
		c.Index(w, r)
	case "/hello":
		c.SayHello(w, r)
	default:
		c.NotFound(w, r)
	}
}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func (c *Controller) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}

func (c *Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	w.Write([]byte("Hello, " + name + "!"))
}
