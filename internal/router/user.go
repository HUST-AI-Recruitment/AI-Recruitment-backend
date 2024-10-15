package router

import (
	"AI-Recruitment-backend/internal/controller"
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.RouterGroup) {
	users := r.Group("/users")

	user := users.Group("/:username")
	user.GET("/profile") // get user profile

	userAuth := users.Group("/:username")
	userAuth.Use(middleware.JwtAuthMiddleware())
	userAuth.PUT("/profile")   // update user profile
	userAuth.POST("/password") // update user password

	auth := r.Group("/auth")
	auth.POST("/register", controller.Register) // register
	auth.POST("/login", controller.Login)       // login
}
