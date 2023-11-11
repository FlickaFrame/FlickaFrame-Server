package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

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
	doerId := jwt.GetUidFromCtx(l.ctx)
	resp = &types.GetVideoCommentResp{}
	comment, err := l.svcCtx.CommentModel.FindOneComment(l.ctx, util.MustString2Int64(req.CommentId))
	if err != nil {
		return nil, err
	}
	resp.Comment, err = NewConvert(l.ctx, l.svcCtx).BuildParentComment(l.ctx, doerId, comment)
	return
}
