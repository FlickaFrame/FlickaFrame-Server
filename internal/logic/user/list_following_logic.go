package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"

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

func (l *ListFollowingLogic) ListFollowing(req *types.ListFollowReq) (resp *types.ListFollowUserResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	contextUserId := req.ContextUserId

	followings, err := l.svcCtx.UserModel.GetUserFollowing(l.ctx, contextUserId, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  req.ListAll,
	})
	if err != nil {
		return nil, err
	}
	// 基本信息
	list, err := NewConvert(l.ctx, l.svcCtx).buildUserBasicInfoList(l.ctx, followings)
	// 互动信息(与当前登录用户)
	info, err := NewConvert(l.ctx, l.svcCtx).BuildUserInteractionInfoList(l.ctx, doerId, followings)
	resp = &types.ListFollowUserResp{
		FollowUser: make([]*types.FollowUser, len(list)),
	}
	for i := range list {
		resp.FollowUser[i] = &types.FollowUser{}
		_ = copier.Copy(resp.FollowUser[i], list[i])
		_ = copier.Copy(resp.FollowUser[i], info[i])
	}
	return
}
