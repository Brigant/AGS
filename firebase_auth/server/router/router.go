package router

import (
	"github.com/Brigant/AGS/firebase_auth/server/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	authController *controller.AuthController
}

func NewRouter(
	authController *controller.AuthController,
) *gin.Engine {
	// Create a new Gin router
	r := gin.Default()

	// Set up the authentication routes
	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)

	return r
}
