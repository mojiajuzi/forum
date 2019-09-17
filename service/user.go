package service

//UserFieldTran 用户验证字段转换
func UserFieldTran() ValidatorFieldTran {
	m := ValidatorFieldTran{}
	m["Name"] = "用户名"
	m["Password"] = "用户密码"
	m["Email"] = "邮箱"
	return m
}
