package oss

import (
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/logic/oss"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUpTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUpTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := oss.NewCreateUpTokenLogic(r.Context(), svcCtx)
		resp, err := l.CreateUpToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
