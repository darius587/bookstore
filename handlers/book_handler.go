package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"bookstore/models"
)

var books []models.Book
var bookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	category := query.Get("category")
	pageStr := query.Get("page")

	page := 1
	limit := 5

	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err == nil && p > 0 {
			page = p
		}
	}

	filtered := []models.Book{}
	for _, b := range books {
		if category == "" || strings.EqualFold(strconv.Itoa(b.CategoryID), category) {
			filtered = append(filtered, b)
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}

	json.NewEncoder(w).Encode(filtered[start:end])
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for _, b := range books {
		if b.ID == id {
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	book.ID = bookID
	bookID++
	books = append(books, book)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	var updated models.Book
	json.NewDecoder(r.Body).Decode(&updated)

	for i, b := range books {
		if b.ID == id {
			if updated.Title != "" {
				b.Title = updated.Title
			}
			if updated.Price > 0 {
				b.Price = updated.Price
			}

			books[i] = b
			json.NewEncoder(w).Encode(b)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, _ := strconv.Atoi(idStr)

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}
