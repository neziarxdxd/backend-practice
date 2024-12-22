package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint      `gorm:"primary_key; auto_increment" json:"user_id"` // Standard field for the primary key
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted bool
}

func CreateUser(tx *gorm.DB) (user *User, err error) {
	user = &User{}
	if err := tx.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetAllUsers(tx *gorm.DB) ([]*User, error) {
	u := []*User{}
	if err := tx.Find(&u).Error; err != nil {
		fmt.Println("Failed to get users")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully get users")
	return u, nil
}
