package user

import (
	"context"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	user_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/jinzhu/copier"
	"strconv"
	"time"

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

// ListFollowing 获取用户关注列表
func (l *ListFollowingLogic) ListFollowing(req *types.ListFollowReq) (resp *types.ListFollowUserResp, err error) {
	doerId, contextUserId := jwt.GetUidFromCtx(l.ctx), req.ContextUserId
	if req.Cursor == 0 {
		req.Cursor = time.Now().UnixMilli()
	}
	followList, err := l.svcCtx.FollowRpc.FollowList(l.ctx, &follow_rpc.FollowListRequest{
		UserId:   contextUserId,
		Cursor:   req.Cursor,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	userIds := make([]int64, 0, len(followList.Items))
	for i := range followList.Items {
		userIds = append(userIds, followList.Items[i].FollowedUserId)
	}
	users, err := l.svcCtx.UserRpc.ListByIds(l.ctx, &user_rpc.ListByIdsRequest{
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.ListFollowUserResp{
		FollowUser: make([]*types.FollowUser, len(users.Users)),
		FeedPagerResp: types.FeedPagerResp{
			Cursor: followList.Cursor,
			IsEnd:  followList.IsEnd,
		},
	}
	for i := range users.Users {
		resp.FollowUser[i] = &types.FollowUser{}
		_ = copier.Copy(resp.FollowUser[i], users.Users[i])
		resp.FollowUser[i].ID = strconv.FormatInt(users.Users[i].Id, 10)
		resp.FollowUser[i].IsFollow = true
		if doerId != contextUserId {
			var follow *follow_rpc.IsFollowResp
			follow, err = l.svcCtx.FollowRpc.IsFollow(l.ctx, &follow_rpc.IsFollowReq{
				UserId:         doerId,
				FollowedUserId: contextUserId,
			})
			if err != nil {
				return nil, err
			}
			resp.FollowUser[i].IsFollow = follow.IsFollow
		}
	}
	return
}
