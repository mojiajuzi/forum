package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/action"
)

func app() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pond"})
	})
	r.POST("/register", action.Register)
	r.POST("/login", action.Login)
	return r
}
