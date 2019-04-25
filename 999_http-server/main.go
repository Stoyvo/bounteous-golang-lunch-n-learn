package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":14715", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path[1:] != "" {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	} else {
		fmt.Fprintf(w, "Hi there!")
	}
}
