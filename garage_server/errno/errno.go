package errno

type Errno interface {
	WithData(data interface{}) Errno
	WithMessage(msg string) Errno
	Error() string
}
type errno struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, msg string) Errno {
	return &errno{
		Code:    code,
		Message: msg,
	}
}

func (e *errno) Error() string {
	return ""
}

func (e *errno) WithData(data interface{}) Errno {
	e.Data = data
	return e
}

func (e *errno) WithMessage(msg string) Errno {
	e.Message = msg
	return e
}

func DecodeError(err error) Errno {
	if err == nil {
		return OK
	}
	switch typed := err.(type) {
	case *errno:
		return typed
	default:
		return InternalServerError.WithMessage(err.Error())
	}
}
