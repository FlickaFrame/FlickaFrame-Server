package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
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

func (l *ListMyFollowingLogic) ListMyFollowing(req *types.ListFollowReq) (resp *types.ListFollowUserResp, err error) {
	req.ContextUserId = jwt.GetUidFromCtx(l.ctx)
	return NewListFollowingLogic(l.ctx, l.svcCtx).ListFollowing(req)
}
