package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckCommentFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckCommentFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckCommentFavoriteLogic {
	return &CheckCommentFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckCommentFavoriteLogic) CheckCommentFavorite(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	resp = &types.FavoriteResp{}
	doerId := jwt.GetUidFromCtx(l.ctx)
	resp.IsFavorite, err = l.svcCtx.FavoriteModel.IsExist(l.ctx, util.MustString2Int64(req.TargetId), doerId)
	return
}
