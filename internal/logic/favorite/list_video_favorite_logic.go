package favorite

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoFavoriteLogic {
	return &ListVideoFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVideoFavoriteLogic) ListVideoFavorite(req *types.ListVideoFavoriteReq) (resp *types.ListVideoFavoriteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
