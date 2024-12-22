package controllers

import (
	"backend-practice/models"

	"github.com/gin-gonic/gin"
)

func (api *API) GetFoods(c *gin.Context) {
	foods, err := models.GetAllFoods(api.db)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, foods)
}

func (api *API) GetFood(c *gin.Context) {
	foodID := c.Param("id")

	food, err := models.GetFood(api.db, foodID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, food)
}

func (api *API) CreateFood(c *gin.Context) {
	food := &models.Food{}
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if _, err := models.CreateFood(api.db, food); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Food created successfully ")
}
