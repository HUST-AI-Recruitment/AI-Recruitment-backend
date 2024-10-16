package router

import (
	"AI-Recruitment-backend/internal/controller"
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.RouterGroup) {
	users := r.Group("/user")
	users.POST("", controller.Register) // register

	user := users.Group("/:username")
	user.GET("/profile") // get user profile

	userAuth := users.Group("/:username")
	userAuth.Use(middleware.JwtAuthMiddleware())
	userAuth.PUT("/profile")  // update user profile
	userAuth.PUT("/password") // update user password

	session := r.Group("/session")
	session.POST("", controller.Login) // login
}
