package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListFollowersLogic {
	return &ListFollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListFollowersLogic) ListFollowers(req *types.ListFollowReq) (resp *types.ListFollowUserResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	followers, err := l.svcCtx.UserModel.GetUserFollowers(l.ctx, req.ContextUserId, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  req.ListAll,
	})
	if err != nil {
		return nil, err
	}
	// 基本信息
	basicUsers, err := NewConvert(l.ctx, l.svcCtx).buildUserBasicInfoList(l.ctx, followers)
	// 互动信息(与当前登录用户)
	info, err := NewConvert(l.ctx, l.svcCtx).BuildUserInteractionInfoList(l.ctx, doerId, followers)
	if err != nil {
		return nil, err
	}
	resp = &types.ListFollowUserResp{
		FollowUser: make([]*types.FollowUser, len(followers)),
	}
	for i := range basicUsers {
		resp.FollowUser[i] = &types.FollowUser{}
		_ = copier.Copy(resp.FollowUser[i], basicUsers[i])
		_ = copier.Copy(resp.FollowUser[i], info[i])
	}
	return
}
