package handlers

import (
	"net/http"
	"strconv"

	"bookstore2/db"
	"bookstore2/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	db.DB.Create(&book)
	c.JSON(http.StatusCreated, book)
}

func GetBooks(c *gin.Context) {
	var books []models.Book

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 5
	offset := (page - 1) * limit

	query := db.DB

	if category := c.Query("category"); category != "" {
		query = query.Where("category_id = ?", category)
	}

	query.Offset(offset).Limit(limit).Find(&books)

	c.JSON(http.StatusOK, books)
}

func GetBookByID(c *gin.Context) {
	var book models.Book

	if err := db.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	if err := db.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	var updated models.Book
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	db.DB.Model(&book).Updates(updated)

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	res := db.DB.Delete(&models.Book{}, c.Param("id"))

	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
