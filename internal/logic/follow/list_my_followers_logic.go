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

type ListMyFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMyFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMyFollowersLogic {
	return &ListMyFollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMyFollowersLogic) ListMyFollowers(req *types.ListMyFollowersReq) (resp *types.ListMyFollowersResp, err error) {
	userID := jwt.GetUidFromCtx(l.ctx)
	followers, err := l.svcCtx.UserModel.GetUserFollowers(l.ctx, userID, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  false,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.ListMyFollowersResp{
		FollowUser: make([]*types.FollowUser, 0, len(followers)),
	}
	for _, follower := range followers {
		followUser := &types.FollowUser{}
		_ = copier.Copy(&follower, followUser)
		resp.FollowUser = append(resp.FollowUser, followUser)
	}
	return
}
