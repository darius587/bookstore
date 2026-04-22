package main

import (
	"bookstore2/db"
	"bookstore2/handlers"
	"bookstore2/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.GET("/books", handlers.GetBooks)
	r.GET("/books/:id", handlers.GetBookByID)

	r.GET("/authors", handlers.GetAuthors)
	r.GET("/categories", handlers.GetCategories)

	r.GET("/books/favorites", middleware.AuthMiddleware(), handlers.GetFavorites)
	r.PUT("/books/:id/favorites", middleware.AuthMiddleware(), handlers.AddFavorite)
	r.DELETE("/books/:id/favorites", middleware.AuthMiddleware(), handlers.RemoveFavorite)

	r.POST("/books", middleware.AuthMiddleware(), handlers.CreateBook)
	r.PUT("/books/:id", middleware.AuthMiddleware(), handlers.UpdateBook)
	r.DELETE("/books/:id", middleware.AuthMiddleware(), handlers.DeleteBook)

	r.POST("/authors", middleware.AuthMiddleware(), handlers.CreateAuthor)
	r.POST("/categories", middleware.AuthMiddleware(), handlers.CreateCategory)

	r.Run(":8080")
}
