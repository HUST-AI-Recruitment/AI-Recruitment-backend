package router

import (
	"AI-Recruitment-backend/internal/controller"
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initUserRouters(r *gin.RouterGroup) {
	users := r.Group("/user")
	users.POST("", controller.Register) // register

	userAuth := users.Group("/:id")
	userAuth.Use(middleware.JwtAuthMiddleware())
	userAuth.GET("/profile", controller.GetUserProfile)    // get user profile
	userAuth.PUT("/profile", controller.UpdateUserProfile) // update user profile
	userAuth.DELETE("")                                    // delete user

	session := r.Group("/session")
	session.POST("", controller.Login) // login
}
