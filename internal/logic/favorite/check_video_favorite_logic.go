package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckVideoFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckVideoFavoriteLogic {
	return &CheckVideoFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckVideoFavoriteLogic) CheckVideoFavorite(req *types.CheckVideoFavoriteReq) (resp *types.CheckVideoFavoriteResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	isFavorite, err := l.svcCtx.FavoriteModel.IsFavoriteVideo(l.ctx, req.VideoId, doerId)
	resp = &types.CheckVideoFavoriteResp{
		IsFavorite: isFavorite,
	}
	return
}
