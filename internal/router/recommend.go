package router

import (
	"AI-Recruitment-backend/internal/controller"
	"AI-Recruitment-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func initRecommendRouters(r *gin.RouterGroup) {
	recommendAuth := r.Group("/recommend")
	recommendAuth.Use(middleware.JwtAuthMiddleware())
	recommendAuth.GET("/jobs", controller.RecommendJobs)               // 获取推荐职位
	recommendAuth.POST("/jobs", controller.RecommendJobsByDescription) // 根据描述获取推荐职位
	recommendAuth.GET("/resumes", controller.RankCandidates)           // 获取推荐简历
}
