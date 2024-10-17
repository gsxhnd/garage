package errno

var (
	OK                  = &errno{0, "OK", nil}
	InternalServerError = &errno{1000, "Internal Server Error", nil}
)
