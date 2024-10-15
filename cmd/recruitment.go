package main

import (
	"AI-Recruitment-backend/internal/global"
	"AI-Recruitment-backend/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	err := initConfig()
	if err != nil {
		log.Fatalf("init.initConfig err: %v", err)
	}

	err = initDB()
	if err != nil {
		log.Fatalf("init.initDB err: %v", err)
	}
}

func main() {
	r := gin.Default()

	router.InitRouters(r)
	r.Run(global.Config.App.Addr)
}
