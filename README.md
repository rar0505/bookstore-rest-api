# Bookstore REST API

This project is a simple REST API in Go.

## Models
- Book
- Author
- Category

## Endpoints

### Books
- GET /books
- POST /books
- GET /books/:id
- PUT /books/:id
- DELETE /books/:id

### Authors
- GET /authors
- POST /authors

### Categories
- GET /categories
- POST /categories

## Features
- CRUD operations
- Pagination
- Filters
- Validation

## Project structure
- main.go
- models/book.go
- models/author.go
- models/category.go
- handlers/book_handler.go
- handlers/author_handler.go
- handlers/category_handler.go
