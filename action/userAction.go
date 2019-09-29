package action

import (
	"net/http"
	"strconv"

	"github.com/mojiajuzi/forum/middleware"

	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/model"
	"github.com/mojiajuzi/forum/service"
	"golang.org/x/crypto/bcrypt"
)

//Register 用户注册
func Register(c *gin.Context) {
	u := model.User{}
	c.BindJSON(&u)
	validate := service.ValidateNew()
	err := validate.Struct(u)
	resp := service.ForumResp{}
	if err != nil {
		errors := service.NewValidatorError(err, service.UserFieldTran())
		resp.Error(http.StatusBadRequest, service.ValidateError, errors)
		c.JSON(500, resp)
		return
	}

	db := model.Db()
	//验证用户是否存在
	oldUser := model.User{}
	db.Where(&model.User{Email: u.Email}).First(&oldUser)
	if oldUser.Email != "" {
		resp.Error(http.StatusBadRequest, "用户已存在，请直接登录", nil)
		c.JSON(500, resp)
		return
	}
	//密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)
	if err != nil {
		resp.Error(http.StatusInternalServerError, "服务异常", nil)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	u.Password = string(hashedPassword)
	//用户存储
	db.Create(&u)
	//发送邮件
	go service.RegisterTemplate(u.Email, u.Name)

	//加密用户数据
	j := middleware.JwtMiddleware()
	uid := strconv.Itoa(int(u.ID))
	token, expire, err := j.TokenGenerator(uid, u)
	if err != nil {
		return
	}
	m := make(map[string]interface{})
	m["expire"] = expire
	m["token"] = token
	resp.Success("注册成功", m)
	c.JSON(200, resp)
	return
}

//User 用户详情
func User(c *gin.Context) {
	//TODO 获取用户详情
}
