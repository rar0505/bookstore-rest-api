package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"bookstore/models"
)

var Categories = []models.Category{}
var CategoryIDCounter = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(category.Name) == "" {
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}

	category.ID = CategoryIDCounter
	CategoryIDCounter++
	Categories = append(Categories, category)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
