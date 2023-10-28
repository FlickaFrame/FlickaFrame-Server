package comment

import (
	"net/http"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/comment"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetVideoCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetVideoCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comment.NewGetVideoCommentLogic(r.Context(), svcCtx)
		resp, err := l.GetVideoComment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
