package controllers

import (
	"backend-practice/models"

	"github.com/gin-gonic/gin"
)

// GetUser displays a message "Hello, World!".
//
// c *gin.Context
func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

// GetUserName retrieves the user ID from the URL parameter and returns a JSON response with a personalized message.
func GetUserName(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Hello, World! " + userID,
	})
}

func CreateUser(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateUser(user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "User created successfully")

}

func GetUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, users)
}
