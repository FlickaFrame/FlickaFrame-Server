package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFavoriteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFavoriteVideoLogic {
	return &UnFavoriteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnFavoriteVideoLogic) UnFavoriteVideo(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.FavoriteModel.DeleteVideoFavorite(l.ctx,
		doerId,
		util.MustString2Int64(req.TargetId))
	return
}
