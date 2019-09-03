package action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/db"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//Register 用户注册
func Register(c *gin.Context) {
	u := db.User{}
	err := c.Bind(&u)
	if err != nil {
		resp := ForumResp{}
		resp.Error(http.StatusBadRequest, err.Error(), nil)
		c.JSON(500, resp)
		return
	}
	fmt.Println(u)
}

//Login 用户登录
func Login(c *gin.Context) {
	//TODO 用户登录
}

//User 用户详情
func User(c *gin.Context) {
	//TODO 获取用户详情
}
