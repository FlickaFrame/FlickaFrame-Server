package follow

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCountFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountFollowersLogic {
	return &CountFollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountFollowersLogic) CountFollowers(req *types.CountFollowReq) (resp *types.CountFollowResp, err error) {
	following, err := l.svcCtx.UserModel.CountFollowing(l.ctx, req.ContextUserId)
	if err != nil {
		return nil, err
	}
	followers, err := l.svcCtx.UserModel.CountFollowers(l.ctx, req.ContextUserId)
	if err != nil {
		return nil, err
	}
	resp = &types.CountFollowResp{
		FollowingCount: following,
		FollowersCount: followers,
	}
	return
}
