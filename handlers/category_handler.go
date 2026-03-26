package handlers

import (
	"encoding/json"
	"net/http"

	"bookstore/models"
)

var categories []models.Category
var categoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil || category.Name == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	category.ID = categoryID
	categoryID++
	categories = append(categories, category)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
