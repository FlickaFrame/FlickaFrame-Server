package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateReplyCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateReplyCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateReplyCommentLogic {
	return &CreateReplyCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateReplyCommentLogic) CreateReplyComment(req *types.CreateReplyCommentReq) (resp *types.CreateReplyCommentResp, err error) {
	return
}
