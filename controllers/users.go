package controllers

import "github.com/gin-gonic/gin"

// displayUsers is a function that displays a message "Hello, World!".
func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func GetUserName(c *gin.Context) {
	// Get the user ID from the URL parameter
	userID := c.Param("id")
	c.JSON(200, gin.H{
		"message": "Hello, World! " + userID,
	})
}
