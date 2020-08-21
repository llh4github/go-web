package common

// Exception 触发异常（panic）。
//
// 主动触发异常，配置中间件可以将异常信息以JSON形式返回。
func Exception(code int, msg string) {
	info := JSONWrapper{Code: code, Msg: msg}
	panic(info)
}

// ExceptionByCode 触发异常（panic）并设定异常代码和文本消息
func ExceptionByCode(code int) {
	c, m := GetErrorInfo(code)
	info := JSONWrapper{Code: c, Msg: m}
	panic(info)
}

const (
	// AuthError 授权异常
	AuthError int = 40100 + iota
	// PwdError 用户或密码错误
	PwdError
)
const (
	// SysError 系统异常
	SysError int = 50000 + iota
	// UnknownError 未知异常
	UnknownError
)

// GetErrorInfo 获取异常信息
//
// 返回异常代码和文本消息。
// TODO 这种方法可能会导致这个方法特别的长。
func GetErrorInfo(errorCode int) (code int, msg string) {
	code = errorCode
	switch errorCode {
	case AuthError:
		msg = "授权异常"
	case PwdError:
		msg = "用户或密码错误"
	case SysError:
		msg = "系统异常"
	default:
		msg = "未知异常"
	}
	return
}
