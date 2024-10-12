package errno

var (
	OK                  = &errno{0, "OK", nil}
	InternalServerError = &errno{100, "Internal Server Error", nil}
)
