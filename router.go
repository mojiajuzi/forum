package main

import (
	"fmt"

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
		auth.Use(middleware.ParseUser())
		auth.GET("/refresh_token", jwt.RefreshHandler)
		auth.GET("/hello", func(c *gin.Context) {
			fmt.Println("hello")
		})
	}
	return r
}
