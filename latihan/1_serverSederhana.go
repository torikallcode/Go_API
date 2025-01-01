package main

import (
	"fmt"
	"log"
	"net/http"
)

func Head(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", Head)
	fmt.Println("Server sedang berjalan di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal()
	}
}
