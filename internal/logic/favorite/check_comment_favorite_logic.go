package favorite

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

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

func (l *CheckCommentFavoriteLogic) CheckCommentFavorite(req *types.CheckCommentFavoriteReq) (resp *types.CheckCommentFavoriteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
