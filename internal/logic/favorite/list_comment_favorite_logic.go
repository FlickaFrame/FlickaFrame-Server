package favorite

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentFavoriteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommentFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentFavoriteLogic {
	return &ListCommentFavoriteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentFavoriteLogic) ListCommentFavorite(req *types.ListCommentFavoriteReq) (resp *types.ListCommentFavoriteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
