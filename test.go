package main

import (
	"io"
	"log"
	"net/http"
)


type hh struct {
}

func (h hh) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>" + r.URL.Path + "</h1>")
}

func main() {
	log.Fatal(http.ListenAndServe(":12345", new(hh)))
}
