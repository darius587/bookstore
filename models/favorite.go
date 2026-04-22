package models

import "time"

type Favorite struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	BookID    uint      `json:"book_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`

	Book Book `gorm:"foreignKey:BookID"`
}
