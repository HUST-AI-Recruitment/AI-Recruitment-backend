package router

import (
	"AI-Recruitment-backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.RouterGroup) {
	users := r.Group("/users")
	user := users.Group("/:username")
	{
		// get user profile
		user.GET("/profile")
		// update user profile
		user.PUT("/profile")
		// update user password
		user.POST("/password")
	}
	auth := r.Group("/auth")
	{
		// register
		auth.POST("/register", controller.Register)
		// login
		auth.POST("/login")
	}
}
