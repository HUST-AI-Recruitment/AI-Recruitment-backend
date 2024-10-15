package router

import (
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initJobRouters(r *gin.RouterGroup) {
	jobs := r.Group("jobs")
	jobs.GET("")     // get all jobs
	jobs.GET("/:id") // get job by id

	jobsAuth := r.Group("jobs")
	jobsAuth.Use(middleware.JwtAuthMiddleware())
	jobsAuth.POST("")       // create jobs
	jobsAuth.PUT("/:id")    // update job by id
	jobsAuth.DELETE("/:id") // delete job by id
}
