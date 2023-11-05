package favorite

import (
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/favorite"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListFavoriteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListFavoriteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := favorite.NewListFavoriteLogic(r.Context(), svcCtx)
		resp, err := l.ListFavorite(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
