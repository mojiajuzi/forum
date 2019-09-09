package action

import (
	"net/http"

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
	if err != nil {
		resp := ForumResp{}
		errors := NewValidatorError(err, u.FieldTrans())
		resp.Error(http.StatusBadRequest, validateError, errors)
		c.JSON(500, resp)
		return
	}

	//验证用户是否存在

	//密码加密

	//用户存储

	//发送邮件
}

//Login 用户登录
func Login(c *gin.Context) {
	//TODO 用户登录
}

//User 用户详情
func User(c *gin.Context) {
	//TODO 获取用户详情
}
