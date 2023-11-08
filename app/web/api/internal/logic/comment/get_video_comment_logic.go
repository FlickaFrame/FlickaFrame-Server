package comment

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoCommentLogic {
	return &GetVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoCommentLogic) GetVideoComment(req *types.GetVideoCommentReq) (resp *types.GetVideoCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
