package models

import (
	"fmt"
	"time"
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

func CreateUser(user *User) error {
	if err := DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]*User, error) {
	u := []*User{}
	if err := DB.Find(&u).Error; err != nil {
		fmt.Println("Failed to get users")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully get users")
	return u, nil
}
