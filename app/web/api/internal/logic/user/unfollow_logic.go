package user

import (
	"context"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnfollowLogic) Unfollow(req *types.UnFollowReq) (resp *types.UnFollowResp, err error) {
	doerUserId := jwt.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.FollowRpc.UnFollow(l.ctx, &follow_rpc.UnFollowRequest{
		UserId:         doerUserId,
		FollowedUserId: req.ContextUserId,
	})
	if err != nil {
		return nil, err
	}
	return
}
