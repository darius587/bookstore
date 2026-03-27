package handlers

import (
	"net/http"
	"strconv"

	"bookstore2/models"

	"github.com/gin-gonic/gin"
)

var books []models.Book
var bookID = 1

func GetBooks(c *gin.Context) {
	category := c.Query("category")
	pageStr := c.Query("page")

	page := 1
	limit := 5

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	filtered := []models.Book{}
	for _, b := range books {
		if category == "" || strconv.Itoa(b.CategoryID) == category {
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

	c.JSON(http.StatusOK, filtered[start:end])
}
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if book.Title == "" || book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	book.ID = bookID
	bookID++
	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}
func GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, b := range books {
		if b.ID == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var updated models.Book
	c.ShouldBindJSON(&updated)

	for i, b := range books {
		if b.ID == id {
			if updated.Title != "" {
				b.Title = updated.Title
			}
			if updated.Price > 0 {
				b.Price = updated.Price
			}

			books[i] = b
			c.JSON(http.StatusOK, b)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
