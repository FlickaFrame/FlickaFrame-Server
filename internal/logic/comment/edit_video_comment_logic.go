package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/pkg/errors"

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
	doer := jwt.GetUidFromCtx(l.ctx)
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.CommentId)
	if err != nil && !errors.Is(err, code.ErrNotFound) {
		return nil, err
	}
	if comment == nil {
		return nil, code.ErrCommentNoExistsError
	}
	if comment.OwnerUID != doer {
		return nil, code.ErrCommentNoPermissionError
	}
	err = l.svcCtx.CommentModel.Update(l.ctx, req.CommentId, req.Content)

	return
}
