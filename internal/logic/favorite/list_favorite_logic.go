package favorite

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFavoriteLogic {
	return &ListFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFavoriteLogic) ListFavorite(req *types.ListFavoriteReq) (resp *types.ListFavoriteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
