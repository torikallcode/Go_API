package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Mhs struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Nim  string `json:"nim"`
}

var mhs = []Mhs{
	{ID: 1, Nama: "akbar", Nim: "101"},
	{ID: 2, Nama: "rikal", Nim: "102"},
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

func updateMhsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	var updateMhs Mhs
	err = json.NewDecoder(r.Body).Decode(&updateMhs)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	for i, mhs_s := range mhs {
		if mhs_s.ID == id {
			mhs[i].Nama = updateMhs.Nama
			mhs[i].Nim = updateMhs.Nim

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&mhs[i])
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
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
	mux.HandleFunc("/mhs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			updateMhsHandler(w, r)
		}
	})

	fmt.Println("Server sedang berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
