package common

// JSONWrapper 统一响应结构体
type JSONWrapper struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// OkResponse 成功处理的响应
func OkResponse(data interface{}) JSONWrapper {
	return JSONWrapper{Code: 200, Msg: "ok", Data: data}
}
