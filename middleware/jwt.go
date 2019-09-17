package middleware

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/appleboy/gin-jwt"

	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/config"
	"github.com/mojiajuzi/forum/model"
	"github.com/mojiajuzi/forum/service"
)

//JwtMiddleware jwt验证中间件
func JwtMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:            "forum",
		Key:              []byte(config.Config("JWT_KEY", "helloginjwt")),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		Authenticator:    authCallback,
		Authorizator:     authPrivCallback,
		Unauthorized:     unAuthFunc,
		TokenLookup:      "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:    "Bearer",
		TimeFunc:         time.Now,
		SigningAlgorithm: "HS256",
		PayloadFunc:      payloadFunc,
	}
}

//验证用户是否存在
func authCallback(c *gin.Context) (interface{}, error) {
	db := model.Db()
	user := model.User{}
	c.BindJSON(&user)
	password := []byte(user.Password)
	record := db.Where("email=?", user.Email).First(&user).RecordNotFound()
	if record {
		err := errors.New("用户名错误")
		return nil, err
	}
	hashPassword := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(hashPassword, password)
	if err != nil {
		err = errors.New("密码错误")
		return nil, err
	}
	return user, nil
}

func payloadFunc(data interface{}) jwt.MapClaims {
	c := jwt.MapClaims{}
	u, ok := data.(model.User)
	if ok {
		c["id"] = u.ID
		c["email"] = u.Email
		c["status"] = u.Status
	}
	return c
}

//验证用户权限
func authPrivCallback(data interface{}, c *gin.Context) bool {

	//TODO 用户权限验证
	return true
}

//验证失败信息回调
func unAuthFunc(c *gin.Context, code int, message string) {
	resp := service.ForumResp{}
	resp.Error(code, message, nil)
	c.JSON(code, resp)
	return
}
