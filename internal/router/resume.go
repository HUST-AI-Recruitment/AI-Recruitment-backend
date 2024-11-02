package router

import (
	"AI-Recruitment-backend/internal/controller"
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initResumeRouters(r *gin.RouterGroup) {
	resumesAuth := r.Group("resumes")
	resumesAuth.Use(middleware.JwtAuthMiddleware())
	resumesAuth.POST("", controller.CreateResume)              // create resumes
	resumesAuth.GET("")                                        // get all resumes
	resumesAuth.GET("/:user_id", controller.GetResumeByUserId) // get resume by id
	resumesAuth.PUT("/:id")                                    // update resume by id
	resumesAuth.DELETE("/:id")                                 // delete resume by id
}
