package router

import "github.com/gin-gonic/gin"

func initResumeRouters(r *gin.RouterGroup) {
	resumes := r.Group("resumes")
	{
		// create resumes
		resumes.POST("")
		// get all resumes
		resumes.GET("")
		// get resume by id
		resumes.GET("/:id")
		// update resume by id
		resumes.PUT("/:id")
		// delete resume by id
		resumes.DELETE("/:id")
	}
}
