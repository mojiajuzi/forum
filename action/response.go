package action

import (
	"strings"

	"github.com/mojiajuzi/forum/db"
	"gopkg.in/go-playground/validator.v9"
)

const (
	validateError = "验证失败"
)

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

//CommonError 错误
type CommonError map[string]interface{}

//NewValidatorError 验证错误处理
func NewValidatorError(err error, m db.ModelFieldTran) CommonError {
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
