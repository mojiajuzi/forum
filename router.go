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

	//分组，添加中间件
	auth := r.Group("/auth")
	website := r.Group("/website")
	jwt := middleware.JwtMiddleware()
	r.POST("/login", jwt.LoginHandler)
	auth.Use(jwt.MiddlewareFunc())
	{
		auth.GET("/refresh_token", jwt.RefreshHandler)
	}

	//用户认证中间件
	website.Use(jwt.MiddlewareFunc(), middleware.ParseUser())
	{
		website.POST("/", action.WebsiteSave)
	}

	return r
}
