package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

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
	doerId := jwt.GetUidFromCtx(l.ctx)
	if !(req.Type == "parent" || req.Type == "child") {
		return nil, code.NoSupportCommentTypeErr
	}
	if req.Type == "parent" {
		err = l.svcCtx.CommentModel.DeleteParentComment(l.ctx, util.MustString2Int64(req.CommentId), doerId)
	} else {
		err = l.svcCtx.CommentModel.DeleteChildComment(l.ctx, util.MustString2Int64(req.CommentId), doerId)
	}
	return nil, err
}
