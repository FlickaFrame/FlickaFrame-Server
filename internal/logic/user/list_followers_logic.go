package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
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

func (l *ListFollowersLogic) ListFollowers(req *types.ListFollowersReq) (resp *types.ListFollowersResp, err error) {
	followers, err := l.svcCtx.UserModel.GetUserFollowers(l.ctx, req.ContextUserId, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  req.ListAll,
	})
	if err != nil {
		return nil, err
	}
	list, err := NewConvert(l.ctx, l.svcCtx).buildUserBasicInfoList(l.ctx, followers)
	resp = &types.ListFollowersResp{
		FollowUser: make([]*types.FollowUser, len(list)),
	}
	err = copier.Copy(&resp.FollowUser, &list)
	return
}
