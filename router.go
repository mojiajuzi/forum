package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/action"
	"github.com/mojiajuzi/forum/middleware"
)

func app() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pond"})
	})
	r.POST("/register", action.Register)
	r.GET("/migrate", action.Migrate)
	auth := r.Group("/auth")
	jwt := middleware.JwtMiddleware()
	r.POST("/login", jwt.LoginHandler)
	auth.Use(jwt.MiddlewareFunc())
	{
		auth.GET("/refresh_token", jwt.RefreshHandler)
	}
	return r
}
