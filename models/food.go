package models

import (
	"time"

	"gorm.io/gorm"
)

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

func GetAllFoods(tx *gorm.DB) ([]*Food, error) {
	f := []*Food{}
	if err := tx.Find(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

func CreateFood(tx *gorm.DB, f *Food) (*Food, error) {
	if err := tx.Create(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

func GetFood(tx *gorm.DB, id string) ([]*Food, error) {
	f := []*Food{}

	if err := tx.Where("food_id = ?", id).Find(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

func UpdateFood(tx *gorm.DB) ([]*Food, error) {
	f := []*Food{}
	if err := tx.Find(&f).Error; err != nil {
		return nil, err

	}
	if err := tx.Save(&f).Error; err != nil {
		return nil, err
	}

	return f, nil
}

func DeleteFood(tx *gorm.DB) ([]*Food, error) {
	f := []*Food{}
	if err := tx.Find(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}
