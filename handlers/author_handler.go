package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var authors []models.Author
var authorID = 1

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil || author.Name == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	author.ID = authorID
	authorID++
	authors = append(authors, author)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(author)
}
