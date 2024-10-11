package errno

var (
	// OK
	OK                  = NewResponse(0, "OK")
	UnknownError        = NewResponse(1, "Unknown Error")
	InternalServerError = &errno{100, "", nil}
	// 服务级错误码
	// ErrServer    = NewError(10001, "服务异常，请联系管理员")
	// ErrParam     = NewError(10002, "参数有误")
	// ErrSignParam = NewError(10003, "签名参数有误")
	// 模块级错误码 - 用户模块
	// ErrUserPhone   = NewError(20101, "用户手机号不合法")
	// ErrUserCaptcha = NewError(20102, "用户验证码有误")
	// ...
)
