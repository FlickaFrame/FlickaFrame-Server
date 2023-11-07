package favorite

import (
	"context"
	favorite_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/favorite"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoLogic {
	return &FavoriteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteVideoLogic) FavoriteVideo(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.FavoriteModel.Create(l.ctx,
		util.MustString2Int64(req.TargetId),
		doerId, favorite_model.VideoFavoriteType)
	return
}
