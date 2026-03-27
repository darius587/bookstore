package handlers

import (
	"net/http"

	"bookstore2/models"

	"github.com/gin-gonic/gin"
)

var categories []models.Category
var categoryID = 1

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil || category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	category.ID = categoryID
	categoryID++
	categories = append(categories, category)

	c.JSON(http.StatusCreated, category)
}
