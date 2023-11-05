package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoCommentLogic {
	return &CreateVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVideoCommentLogic) CreateVideoComment(req *types.CreateVideoCommentReq) (resp *types.CreateVideoCommentResp, err error) {
	doer := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.CommentModel.CreateParentComment(l.ctx, doer, req.VideoId, req.Content)
	return
}
