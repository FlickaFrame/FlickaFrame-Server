package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

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
	doer := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.CommentModel.CreateReplyComment(l.ctx, doer, req.VideoId, req.Content, req.ParentId, req.TargetId)
	return
}
