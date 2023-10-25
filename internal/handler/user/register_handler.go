package user

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/user"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Debug(err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		if err := svcCtx.Validate.Struct(&req); err != nil {
			logx.Debug(err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
