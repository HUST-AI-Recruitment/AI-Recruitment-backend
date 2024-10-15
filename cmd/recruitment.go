package main

import (
	"AI-Recruitment-backend/internal/config"
	"AI-Recruitment-backend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouters(r)
	r.Run(config.C.App.Addr)
}
