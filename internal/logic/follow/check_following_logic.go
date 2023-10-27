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
	resp.Status = l.svcCtx.UserModel.IsFollowing(l.ctx, req.DoerUserId, req.ContextUserId)
	return
}
