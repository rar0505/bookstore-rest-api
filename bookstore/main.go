package main

import (
	"fmt"
	"net/http"

	"bookstore/handlers"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetBooks(w, r)
	case http.MethodPost:
		handlers.CreateBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetBookByID(w, r)
	case http.MethodPut:
		handlers.UpdateBook(w, r)
	case http.MethodDelete:
		handlers.DeleteBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func authorsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetAuthors(w, r)
	case http.MethodPost:
		handlers.CreateAuthor(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handlers.GetCategories(w, r)
	case http.MethodPost:
		handlers.CreateCategory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/books", booksHandler)
	http.HandleFunc("/books/", bookByIDHandler)

	http.HandleFunc("/authors", authorsHandler)
	http.HandleFunc("/categories", categoriesHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
