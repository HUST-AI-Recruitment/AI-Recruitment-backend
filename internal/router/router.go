package router

import (
	"AI-Recruitment-backend/internal/global"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	api := r.Group(global.Config.App.ApiPrefix)
	initUserRouters(api)
	initJobRouters(api)
	initResumeRouters(api)
	initApplicationRouters(api)
}
