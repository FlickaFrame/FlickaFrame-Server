package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowingLogic {
	return &ListFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowingLogic) ListFollowing(req *types.ListFollowingReq) (resp *types.ListFollowingResp, err error) {
	followings, err := l.svcCtx.UserModel.GetUserFollowing(l.ctx, req.ContextUserId, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  false,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ListFollowingResp{
		FollowUser: make([]*types.UserBasicInfo, 0, len(followings)),
	}
	for _, follower := range followings {
		followUser := &types.UserBasicInfo{}
		_ = copier.Copy(&follower, followUser)
		resp.FollowUser = append(resp.FollowUser, followUser)
	}
	return
}
