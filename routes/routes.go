package routes

import (
	"backend-practice/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	users := r.Group("/users")
	{
		users.GET("/", controllers.GetUser)
		users.GET("/:id", controllers.GetUserName)
		// users.GET("/:id", controllers.GetUser)
		// users.POST("/", controllers.CreateUser)
		// users.PATCH("/:id", controllers.UpdateUser)
		// users.DELETE("/:id", controllers.DeleteUser)
	}

	return r
}
