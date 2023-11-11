package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFavoriteCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnFavoriteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFavoriteCommentLogic {
	return &UnFavoriteCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnFavoriteCommentLogic) UnFavoriteComment(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	resp = &types.FavoriteResp{IsFavorite: false}
	doerId := jwt.GetUidFromCtx(l.ctx)
	// redis desc
	count, err := l.svcCtx.BizRedis.Decr(cacheCommentLikeCount(util.MustString2Int64(req.TargetId)))
	resp.LikeCount = int(count)
	err = l.svcCtx.FavoriteModel.DeleteCommentFavorite(l.ctx,
		doerId,
		util.MustString2Int64(req.TargetId))
	return
}
