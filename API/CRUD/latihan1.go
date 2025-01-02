package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Mhs struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Nim  string `json:"nim"`
}

var mhs = []Mhs{
	{ID: 1, Nama: "Torikal", Nim: "201"},
	{ID: 2, Nama: "Akbar", Nim: "202"},
}

func getMhs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mhs)
}

func createMhs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newMhs Mhs
	err := json.NewDecoder(r.Body).Decode(&newMhs)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newMhs.ID = len(mhs) + 1
	mhs = append(mhs, newMhs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMhs)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/mhs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getMhs(w, r)
		} else if r.Method == http.MethodPost {
			createMhs(w, r)
		}
	})

	fmt.Println("Server sedang berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
