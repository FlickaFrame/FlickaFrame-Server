package follow

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"

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

func (l *ListMyFollowingLogic) ListMyFollowing(req *types.ListMyFollowingReq) (resp *types.ListMyFollowingResp, err error) {
	userId := jwt.GetUidFromCtx(l.ctx)
	followings, err := l.svcCtx.UserModel.GetUserFollowing(l.ctx, userId, orm.ListOptions{
		PageSize: req.Limit,
		Page:     req.Page,
		ListAll:  false,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.ListMyFollowingResp{
		FollowUser: make([]*types.FollowUser, 0, len(followings)),
	}
	for _, follower := range followings {
		followUser := &types.FollowUser{}
		_ = copier.Copy(&follower, followUser)
		resp.FollowUser = append(resp.FollowUser, followUser)
	}
	return
}
