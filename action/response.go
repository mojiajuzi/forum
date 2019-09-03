package action

//ForumResp 响应结构体
type ForumResp struct {
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	ErrCode int         `json:"err_code"`
}

//Success 成功响应
func (f *ForumResp) Success(msg string, content interface{}) {
	f.Status = true
	f.ErrCode = 0
	f.Msg = msg
	f.Data = content
}

//Error 失败响应
func (f *ForumResp) Error(code int, msg string, content interface{}) {
	f.Status = false
	f.ErrCode = code
	f.Msg = msg
	f.Data = content
}
