package router

import (
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initJobRouters(r *gin.RouterGroup) {
	jobsAuth := r.Group("jobs")
	jobsAuth.Use(middleware.JwtAuthMiddleware())
	jobsAuth.GET("")        // get all jobs
	jobsAuth.GET("/:id")    // get job by id
	jobsAuth.POST("")       // create jobs
	jobsAuth.PUT("/:id")    // update job by id
	jobsAuth.DELETE("/:id") // delete job by id
}
