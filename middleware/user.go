package middleware

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/service"
)

//ParseUser 解析用户标识
func ParseUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := jwt.ExtractClaims(c)
		for a, index := range t {
			c.Set(a, index)
		}

		//判断用户当前状态
		resp := service.ForumResp{}
		if status, ok := c.Get("status"); ok {
			if s, ok := status.(int); ok && s == 0 {
				resp.Error(http.StatusForbidden, "账户禁用", nil)
				c.JSON(403, resp)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
