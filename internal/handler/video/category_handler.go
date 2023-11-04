package video

import (
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := video.NewCategoryLogic(r.Context(), svcCtx)
		resp, err := l.Category()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
