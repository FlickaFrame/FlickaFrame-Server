package comment

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVideoCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoCommentsLogic {
	return &ListVideoCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVideoCommentsLogic) ListVideoComments(req *types.ListVideoCommentsReq) (resp *types.ListVideoCommentsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
