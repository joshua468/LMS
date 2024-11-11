package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/lms-backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User routes
	r.GET("/users/:userID", controllers.GetUserProfile)

	return r
}
