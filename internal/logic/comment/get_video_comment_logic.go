package comment

import (
	"context"
	"github.com/jinzhu/copier"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

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
	resp = &types.GetVideoCommentResp{}
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, req.CommentId)
	if err != nil {
		return nil, err
	}
	var commentRsp types.Commnent
	_ = copier.Copy(&commentRsp, comment)
	resp.Commnent = &commentRsp
	return
}
