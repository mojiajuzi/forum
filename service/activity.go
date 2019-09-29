package service

//ActivityFieldTran 用户验证字段转换
func ActivityFieldTran() ValidatorFieldTran {
	m := ValidatorFieldTran{}
	m["Title"] = "活动标题"
	m["Cover"] = "封面"
	m["Content"] = "内容"
	m["StartAt"] = "开始时间"
	m["EndAt"] = "结束时间"
	m["WillTotal"] = "票数"
	return m
}
