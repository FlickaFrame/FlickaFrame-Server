package tag

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RankLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRankLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RankLogic {
	return &RankLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RankLogic) Rank(req *types.TagReq) (resp *types.FollowResp, err error) {
	// todo: add your logic here and delete this line

	return
}
