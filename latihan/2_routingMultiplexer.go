package main

import (
	"fmt"
	"log"
	"net/http"
)

func Init(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my website")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the about page")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contact us at torikal@akbar.com")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Init)
	mux.HandleFunc("/about", About)
	mux.HandleFunc("/contact", Contact)

	fmt.Println("Server sedang berjalan di localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal()
	}
}
