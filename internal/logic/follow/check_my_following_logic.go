package follow

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

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
	doerUserId := jwt.GetUidFromCtx(l.ctx)
	resp.Status = l.svcCtx.UserModel.IsFollowing(l.ctx, doerUserId, req.ContextUserId)
	return
}
