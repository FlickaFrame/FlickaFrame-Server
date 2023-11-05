package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFavoriteLogic {
	return &UnFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnFavoriteLogic) UnFavorite(req *types.UnFavoriteReq) (resp *types.UnFavoriteResp, err error) {
	return nil, l.svcCtx.FavoriteModel.Delete(l.ctx, jwt.GetUidFromCtx(l.ctx), req.TargetId)
}
