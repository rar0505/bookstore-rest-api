package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"bookstore/models"
)

var Books = []models.Book{}
var BookIDCounter = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryFilter := r.URL.Query().Get("category_id")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 5

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err == nil && l > 0 {
			limit = l
		}
	}

	filteredBooks := []models.Book{}

	if categoryFilter != "" {
		categoryID, err := strconv.Atoi(categoryFilter)
		if err == nil {
			for _, book := range Books {
				if book.CategoryID == categoryID {
					filteredBooks = append(filteredBooks, book)
				}
			}
		}
	} else {
		filteredBooks = Books
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(filteredBooks) {
		start = len(filteredBooks)
	}
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	json.NewEncoder(w).Encode(filteredBooks[start:end])
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(book.Title) == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if book.AuthorID <= 0 {
		http.Error(w, "AuthorID is required", http.StatusBadRequest)
		return
	}

	if book.CategoryID <= 0 {
		http.Error(w, "CategoryID is required", http.StatusBadRequest)
		return
	}

	if book.Price < 0 {
		http.Error(w, "Price must be greater than or equal to 0", http.StatusBadRequest)
		return
	}

	book.ID = BookIDCounter
	BookIDCounter++
	Books = append(Books, book)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for _, book := range Books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook models.Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(updatedBook.Title) == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	if updatedBook.AuthorID <= 0 {
		http.Error(w, "AuthorID is required", http.StatusBadRequest)
		return
	}

	if updatedBook.CategoryID <= 0 {
		http.Error(w, "CategoryID is required", http.StatusBadRequest)
		return
	}

	if updatedBook.Price < 0 {
		http.Error(w, "Price must be greater than or equal to 0", http.StatusBadRequest)
		return
	}

	for i, book := range Books {
		if book.ID == id {
			updatedBook.ID = id
			Books[i] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Book deleted successfully",
			})
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
