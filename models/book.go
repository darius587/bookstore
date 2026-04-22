package models

type Book struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Title      string  `json:"title"`
	AuthorID   uint    `json:"author_id"`
	CategoryID uint    `json:"category_id"`
	Price      float64 `json:"price"`
}
