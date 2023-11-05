package xcode

import "github.com/FlickaFrame/FlickaFrame-Server/pkg/container"

var (
	OK                 = add(0, "OK")
	NoLogin            = add(101, "NOT_LOGIN")
	RequestErr         = add(400, "INVALID_ARGUMENT")
	Unauthorized       = add(401, "未登录!")
	AccessDenied       = add(403, "PERMISSION_DENIED")
	NotFound           = add(404, "NOT_FOUND")
	MethodNotAllowed   = add(405, "METHOD_NOT_ALLOWED")
	Canceled           = add(498, "CANCELED")
	ServerErr          = add(500, "INTERNAL_ERROR")
	ServiceUnavailable = add(503, "UNAVAILABLE")
	Deadline           = add(504, "DEADLINE_EXCEEDED")
	LimitExceed        = add(509, "RESOURCE_EXHAUSTED")
	DB_ERROR           = add(510, "DB_ERROR")
)

var errSet = make(container.Set[string])

func Add(code int, msg string) Code {
	return add(code, msg)
}

func add(code int, msg string) Code {
	if errSet.Contains(msg) {
		panic("duplicate error code: " + msg)
	}
	return Code{code: code, msg: msg}
}
