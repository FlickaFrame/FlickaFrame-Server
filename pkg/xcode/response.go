package xcode

import (
	"context"
	"net/http"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return http.StatusOK, Status{
		Code:    int32(code.Code()),
		Message: err.Error(),
		Success: false,
	}
}

func OkHandler(ctx context.Context, v any) any {
	return Status{
		Code:    0,
		Message: "ok",
		Success: true,
		Data:    v,
	}
}
