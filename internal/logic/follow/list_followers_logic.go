package follow

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

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
		PageSize: req.Limit,
		Page:     req.Page,
		ListAll:  false,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.ListFollowersResp{
		FollowUser: make([]*types.FollowUser, 0, len(followers)),
	}
	for _, follower := range followers {
		followUser := &types.FollowUser{}
		_ = copier.Copy(&follower, followUser)
		resp.FollowUser = append(resp.FollowUser, followUser)
	}
	return
}
