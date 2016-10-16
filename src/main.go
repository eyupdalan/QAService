package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Message struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome %s", r.URL.Path[1:])
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to the my API"}
	b, err := json.Marshal(m)

	if (err != nil) {
		panic(err)
	}

	w.Write(b)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", handler)
	http.ListenAndServe(":80", nil)
}