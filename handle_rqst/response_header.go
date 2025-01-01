package main

import (
	"fmt"
	"net/http"
)

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello world")
}

func main() {
	http.HandleFunc("/", htmlVsPlain)
	http.ListenAndServe("", nil)
}
