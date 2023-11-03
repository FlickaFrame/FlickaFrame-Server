package user

import (
	"context"
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

func (l *ListFollowingLogic) ListFollowing(req *types.ListFollowingReq) (resp *types.ListFollowingResp, err error) {
	followings, err := l.svcCtx.UserModel.GetUserFollowing(l.ctx, req.ContextUserId, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  req.ListAll,
	})
	if err != nil {
		return nil, err
	}
	list, err := NewConvert(l.ctx, l.svcCtx).buildUserBasicInfoList(l.ctx, followings)
	resp = &types.ListFollowingResp{
		FollowUser: make([]*types.FollowUser, len(list)),
	}
	err = copier.Copy(&resp.FollowUser, &list)
	return
}
