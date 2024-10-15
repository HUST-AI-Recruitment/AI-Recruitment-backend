package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}
