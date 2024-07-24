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
		users.GET("/", controllers.GetUsers)
		users.GET("/:id", controllers.GetUserName)
		users.POST("/", controllers.CreateUser)

		// users.GET("/:id", controllers.GetUser)
		// users.POST("/", controllers.CreateUser)
		// users.PATCH("/:id", controllers.UpdateUser)
		// users.DELETE("/:id", controllers.DeleteUser)
	}

	return r
}
