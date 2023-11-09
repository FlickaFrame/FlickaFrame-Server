package user

import (
	"context"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	user_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
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
	list, err := l.svcCtx.FollowRpc.FollowList(l.ctx, &follow_rpc.FollowListRequest{
		UserId:   req.ContextUserId,
		Cursor:   0,
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	userIds := make([]int64, 0, len(list.Items))
	for i := range list.Items {
		userIds = append(userIds, list.Items[i].FollowedUserId)
	}
	users, err := l.svcCtx.UserRpc.ListByIds(l.ctx, &user_rpc.ListByIdsRequest{
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.ListFollowUserResp{
		FollowUser: make([]*types.FollowUser, len(users.Users)),
		Total:      int64(users.Total),
	}
	for i := range users.Users {
		resp.FollowUser[i] = &types.FollowUser{}
		_ = copier.Copy(resp.FollowUser[i], users.Users[i])
		var follow *follow_rpc.IsFollowResp
		follow, err = l.svcCtx.FollowRpc.IsFollow(l.ctx, &follow_rpc.IsFollowReq{
			UserId:         doerId,
			FollowedUserId: req.ContextUserId,
		})
		if err != nil {
			return nil, err
		}
		resp.FollowUser[i].IsFollow = follow.IsFollow
	}
	return
}
