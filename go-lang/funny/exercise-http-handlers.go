package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s", s.Greeting, s.Who)
}

func main() {
	http.Handle("/string", String("I am lidachao"))
	http.Handle("/struct", &Struct{"hello", "lidachao"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
