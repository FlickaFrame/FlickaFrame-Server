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

type DeleteVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoCommentLogic {
	return &DeleteVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoCommentLogic) DeleteVideoComment(req *types.DeleteVideoCommentReq) (resp *types.DeleteVideoCommentResp, err error) {
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
	err = l.svcCtx.CommentModel.Delete(l.ctx, req.CommentId)

	return
}
