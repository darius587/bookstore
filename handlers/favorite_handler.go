package handlers

import (
	"net/http"
	"strconv"

	"bookstore2/db"
	"bookstore2/models"

	"github.com/gin-gonic/gin"
)

func AddFavorite(c *gin.Context) {
	userID := c.GetUint("user_id")

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}

	var existing models.Favorite
	err = db.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		First(&existing).Error

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already in favorites"})
		return
	}

	favorite := models.Favorite{
		UserID: userID,
		BookID: uint(bookID),
	}

	if err := db.DB.Create(&favorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "added to favorites"})
}

func GetFavorites(c *gin.Context) {
	userID := c.GetUint("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 5
	offset := (page - 1) * limit

	var books []models.Book

	err := db.DB.
		Table("books").
		Joins("JOIN favorites ON favorites.book_id = books.id").
		Where("favorites.user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&books).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func RemoveFavorite(c *gin.Context) {
	userID := c.GetUint("user_id")

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book id"})
		return
	}

	result := db.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Delete(&models.Favorite{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "removed from favorites"})
}
