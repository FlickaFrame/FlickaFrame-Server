package xcode

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)
	var success bool
	return http.StatusOK, Status{
		Code:    int32(code.Code()),
		Message: err.Error(),
		Success: &success,
	}
}

func OkHandler(ctx context.Context, v any) any {
	success := true
	return Status{
		Code:    0,
		Message: "ok",
		Success: &success,
		Data:    v,
	}
}

// UnAuthorizedCallback 统一未授权响应格式
func UnAuthorizedCallback(w http.ResponseWriter, r *http.Request, err error) {
	httpx.Error(w, Unauthorized)
}
