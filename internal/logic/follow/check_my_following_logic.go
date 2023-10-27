package follow

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckMyFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckMyFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckMyFollowingLogic {
	return &CheckMyFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckMyFollowingLogic) CheckMyFollowing(req *types.CheckMyFollowingReq) (resp *types.CheckMyFollowingResp, err error) {
	// todo: add your logic here and delete this line

	return
}
