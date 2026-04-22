package handlers

import (
	"net/http"

	"bookstore2/db"
	"bookstore2/models"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	db.DB.Create(&category)
	c.JSON(http.StatusCreated, category)
}

func GetCategories(c *gin.Context) {
	var categories []models.Category

	db.DB.Find(&categories)

	c.JSON(http.StatusOK, categories)
}
