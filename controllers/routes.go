package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/henvo/golang-gin-gorm-starter/models"
)

// SetupRouter sets up the router.
func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := &API{handler: r, db: models.DB}
	food := r.Group("/food")
	{
		food.GET("/:id", api.GetFood)
		food.GET("/", api.GetFoods)
	}

	return r
}
