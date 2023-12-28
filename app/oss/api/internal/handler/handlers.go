package handler

import (
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/types"
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/logic"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EndpointHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewEndpointLogic(r.Context(), ctx)
		resp, err := l.Endpoint()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

func CreateUpTokenHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUpTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUpTokenLogic(r.Context(), ctx)
		resp, err := l.CreateUpToken(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
