package comment

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditVideoCommentLogic {
	return &EditVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditVideoCommentLogic) EditVideoComment(req *types.EditVideoCommentReq) (resp *types.EditVideoCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
