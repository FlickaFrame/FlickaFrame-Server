package follow

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckFollowingLogic {
	return &CheckFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckFollowingLogic) CheckFollowing(req *types.CheckFollowingReq) (resp *types.CheckFollowingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
