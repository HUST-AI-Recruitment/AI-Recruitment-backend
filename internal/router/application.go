package router

import (
	"AI-Recruitment-backend/internal/controller"
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initApplicationRouters(r *gin.RouterGroup) {
	applicationsAuth := r.Group("applications")
	applicationsAuth.Use(middleware.JwtAuthMiddleware())
	applicationsAuth.POST("", controller.CreateApplication)                 // apply for job
	applicationsAuth.GET("", controller.GetApplicationsByUserId)            // get user's applications
	applicationsAuth.GET("/job/:job_id", controller.GetApplicationsByJobId) // get applications by job
	applicationsAuth.PUT("/:id")                                            // update progress of application
}
