package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFavoriteLogic {
	return &CheckFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckFavoriteLogic) CheckFavorite(req *types.CheckFavoriteReq) (resp *types.CheckFavoriteResp, err error) {
	err = l.svcCtx.FavoriteModel.IsExist(l.ctx, req.TargetId, jwt.GetUidFromCtx(l.ctx))
	resp = &types.CheckFavoriteResp{
		IsFavorite: err == nil,
	}
	return
}
