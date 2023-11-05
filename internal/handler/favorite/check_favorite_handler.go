package favorite

import (
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/favorite"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckFavoriteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckFavoriteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := favorite.NewCheckFavoriteLogic(r.Context(), svcCtx)
		resp, err := l.CheckFavorite(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
