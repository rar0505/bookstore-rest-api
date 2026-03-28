package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"bookstore/models"
)

var Authors = []models.Author{}
var AuthorIDCounter = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var author models.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(author.Name) == "" {
		http.Error(w, "Author name is required", http.StatusBadRequest)
		return
	}

	author.ID = AuthorIDCounter
	AuthorIDCounter++
	Authors = append(Authors, author)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}
