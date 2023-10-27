package follow

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMyFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMyFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMyFollowingLogic {
	return &ListMyFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMyFollowingLogic) ListMyFollowing(req *types.ListMyFollowingReq) (resp *types.ListMyFollowingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
