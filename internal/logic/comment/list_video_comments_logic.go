package comment

import (
	"context"
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/comment"
	"github.com/jinzhu/copier"

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
	resp = &types.ListVideoCommentsResp{}
	comments, err := l.svcCtx.CommentModel.List(l.ctx, comment_model.Option{
		VideoID: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	var commentsRsp []*types.Commnent
	for _, comment := range comments {
		var commentRsp types.Commnent
		_ = copier.Copy(&commentRsp, comment)
		commentsRsp = append(commentsRsp, &commentRsp)
	}
	resp.Comments = commentsRsp
	return
}
