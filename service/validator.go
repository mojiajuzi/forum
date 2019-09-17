package service

import (
	"strings"

	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

//ValidatorFieldTran 字段验证名称转换
type ValidatorFieldTran map[string]string

var (
	validate *validator.Validate
	trans    ut.Translator
)

const (
	ValidateError = "验证失败"
)

func init() {
	zh := zhongwen.New()
	uni := ut.New(zh, zh)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

//ValidateNew 校验数据
func ValidateNew() *validator.Validate {
	return validate
}

//CommonError 错误
type CommonError map[string]interface{}

//ForumResp 响应结构体
type ForumResp struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	ErrCode int         `json:"err_code"`
	Errors  CommonError `json:"errors"`
}

//Success 成功响应
func (f *ForumResp) Success(msg string, content interface{}) {
	f.Status = true
	f.ErrCode = 0
	f.Msg = msg
	f.Data = content
}

//Error 失败响应
func (f *ForumResp) Error(code int, msg string, err CommonError) {
	f.Status = false
	f.ErrCode = code
	f.Msg = msg
	f.Errors = err
}

//NewValidatorError 验证错误处理
func NewValidatorError(err error, m ValidatorFieldTran) CommonError {
	res := CommonError{}
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		transtr := e.Translate(trans)
		f := strings.ToLower(e.Field())
		if rp, ok := m[e.Field()]; ok {
			res[f] = strings.Replace(transtr, e.Field(), rp, 1)
		} else {
			res[f] = transtr
		}
	}
	return res
}
