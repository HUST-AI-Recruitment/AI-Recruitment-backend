package router

import (
	"AI-Recruitment-backend/internal/config"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group(config.C.App.ApiPrefix)
	initUserRouters(api)
	initJobRouters(api)
	initResumeRouters(api)
}
