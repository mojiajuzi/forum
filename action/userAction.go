package action

import (
	"net/http"

	"github.com/mojiajuzi/forum/service"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/model"

	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	validator "gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

//Register 用户注册
func Register(c *gin.Context) {
	u := model.User{}
	c.BindJSON(&u)
	err := validate.Struct(u)
	resp := ForumResp{}
	if err != nil {
		errors := NewValidatorError(err, u.FieldTrans())
		resp.Error(http.StatusBadRequest, validateError, errors)
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
	u.Password = ""
	resp.Success("注册成功", u)
	c.JSON(200, resp)
	return
}

//User 用户详情
func User(c *gin.Context) {
	//TODO 获取用户详情
}
