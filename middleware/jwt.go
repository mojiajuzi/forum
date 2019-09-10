package middleware

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/action"
	"github.com/mojiajuzi/forum/model"
	"github.com/mojiajuzi/forum/service"
)

//JwtMiddleware jwt验证中间件
func JwtMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "zone name just for test",
		Key:           []byte(service.Config("JWT_KEY", "helloginjwt")),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: authCallback,
		Authorizator:  authPrivCallback,
		Unauthorized:  unAuthFunc,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

//验证用户是否存在
func authCallback(c *gin.Context) (interface{}, error) {
	password := c.PostForm("password")
	userID := c.PostForm("email")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, err
	}

	db := model.Db()
	user := model.User{}
	db.Where("email=? AND password=?", userID, hashedPassword).First(&user)
	if user.Email == "" {
		err = errors.New("用户名或密码错误")
		return nil, err
	}
	return user, nil
}

//验证用户权限
func authPrivCallback(data interface{}, c *gin.Context) bool {

	//TODO 用户权限验证
	return true
}

//验证失败信息回调
func unAuthFunc(c *gin.Context, code int, message string) {
	resp := action.ForumResp{}
	resp.Error(code, message, nil)
	c.JSON(code, resp)
	return
}
