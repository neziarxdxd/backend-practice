package models

import "time"

type Food struct {
	FoodID    uint    `gorm:"primary_key; auto_increment" json:"food_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Store     *string `json:"store"`
	Link      *string `json:"link"`
	Notes     *string `json:"notes"`
	IsDeleted bool    `json:"is_deleted"`
	UserID    uint    `json:"user_id"`
	User      User
	CreatedAt time.Time `json:"created_at"`
}
