package service

import (
	"github.com/joho/godotenv"
)

var myEnv map[string]string

func init() {
	env, err := godotenv.Read()
	if err != nil {
		panic("获取配置文件错误")
	}
	myEnv = env
}

//Config 获取所有配置
//name 参数名称
//d 默认值
func Config(name, d string) string {
	if v, ok := myEnv[name]; ok {
		return v
	}
	return d
}
