package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

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
	resp = &types.FavoriteResp{IsFavorite: true}
	doerId := jwt.GetUidFromCtx(l.ctx)
	// 检查视频是否存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, util.MustString2Int64(req.TargetId))
	if err != nil {
		logx.Info(err)
		return nil, code.VideoNotExistError
	}
	err = l.svcCtx.FavoriteModel.CreateVideoFavorite(l.ctx,
		doerId,
		util.MustString2Int64(req.TargetId))
	if err != nil {
		logx.Info(err)
		return nil, code.DuplicateFavoriteErr
	}
	return
}
