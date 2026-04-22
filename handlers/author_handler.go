package handlers

import (
	"net/http"

	"bookstore2/db"
	"bookstore2/models"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	db.DB.Create(&author)
	c.JSON(http.StatusCreated, author)
}

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	db.DB.Find(&authors)

	c.JSON(http.StatusOK, authors)
}
