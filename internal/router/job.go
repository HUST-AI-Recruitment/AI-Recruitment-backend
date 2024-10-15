package router

import "github.com/gin-gonic/gin"

func initJobRouters(r *gin.RouterGroup) {
	jobs := r.Group("jobs")
	{
		// create jobs
		jobs.POST("")
		// get all jobs
		jobs.GET("")
		// get job by id
		jobs.GET("/:id")
		// update job by id
		jobs.PUT("/:id")
		// delete job by id
		jobs.DELETE("/:id")
	}
}
