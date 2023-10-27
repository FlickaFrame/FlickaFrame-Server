package follow

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMyFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMyFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMyFollowersLogic {
	return &ListMyFollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMyFollowersLogic) ListMyFollowers(req *types.ListMyFollowersReq) (resp *types.ListMyFollowersResp, err error) {
	// todo: add your logic here and delete this line

	return
}
