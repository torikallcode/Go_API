package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "M. Torikal", Email: "Torikal@gmail.com"},
	{ID: 2, Name: "Akbar", Email: "Akbar@gmail.com"},
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", getUserHandler)

	fmt.Println("Server sedang berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
