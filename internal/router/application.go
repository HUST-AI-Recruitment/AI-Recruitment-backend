package router

import (
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initApplicationRouters(r *gin.RouterGroup) {
	applicationsAuth := r.Group("applications")
	applicationsAuth.Use(middleware.JwtAuthMiddleware())
	applicationsAuth.POST("")    // apply for job
	applicationsAuth.GET("")     // get all applications
	applicationsAuth.PUT("/:id") // update progress of application
}
