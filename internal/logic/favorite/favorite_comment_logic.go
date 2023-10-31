package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteCommentLogic {
	return &FavoriteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteCommentLogic) FavoriteComment(req *types.FavoriteCommentReq) (resp *types.FavoriteCommentResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.FavoriteModel.FavoriteComment(l.ctx, doerId, req.VideoId, req.IsFavorite)
	return
}
