package action

import (
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//Register 用户注册
func Register(w http.ResponseWriter, r *http.Request) {
	//TODO 用户注册
}

//Login 用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	//TODO 用户登录
}

//User 用户详情
func User(w http.ResponseWriter, r *http.Request) {
	//TODO 获取用户详情
}
