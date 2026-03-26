package main

import (
	"fmt"
	"net/http"

	"bookstore/handlers"
)

func main() {

	http.HandleFunc("/books", handlers.GetBooks)
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetBookByID(w, r)
		case http.MethodPut:
			handlers.UpdateBook(w, r)
		case http.MethodDelete:
			handlers.DeleteBook(w, r)
		}
	})
	http.HandleFunc("/books/create", handlers.CreateBook)

	http.HandleFunc("/authors", handlers.GetAuthors)
	http.HandleFunc("/authors/create", handlers.CreateAuthor)
	
	http.HandleFunc("/categories", handlers.GetCategories)
	http.HandleFunc("/categories/create", handlers.CreateCategory)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
