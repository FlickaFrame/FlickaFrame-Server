package logic

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/oss"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/pb/user_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByIdsLogic {
	return &ListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListByIdsLogic) ListByIds(in *user_service.ListByIdsRequest) (*user_service.ListByIdsResponse, error) {
	users, err := l.svcCtx.UserModel.FindByIDs(l.ctx, in.UserIds)
	if err != nil {
		return nil, err
	}
	UserInfos := make([]*user_service.UserInfoResponse, 0, len(users))
	for _, user := range users {
		// 用户的头像地址实际上是一个oss的key，需要通过oss服务获取访问地址
		url, err := l.svcCtx.OssRpc.GetFileAccessUrl(l.ctx, &oss.GetFileAccessUrlRequest{Key: user.AvatarUrl})
		if err != nil {
			return nil, err
		}
		// 用户背景图片地址实际上是一个oss的key，需要通过oss服务获取访问地址
		bgUrl, err := l.svcCtx.OssRpc.GetFileAccessUrl(l.ctx, &oss.GetFileAccessUrlRequest{Key: user.BackgroundUrl})
		if err != nil {
			return nil, err
		}
		user.AvatarUrl = url.Url
		userItems := &user_service.UserInfoResponse{
			Id:             user.ID,
			NickName:       user.NickName,
			AvatarUrl:      user.AvatarUrl,
			Slogan:         user.Slogan,
			Gender:         user.Gender,
			Age:            user.Age,
			Mobile:         user.Phone,
			FollowingCount: user.FollowingCount,
			FollowerCount:  user.FollowerCount,
			LikeCount:      0, //TODO: 获得点赞数
			CollectCount:   0, //TODO: 收藏数
			BackgroundUrl:  bgUrl.Url,
		}
		UserInfos = append(UserInfos, userItems)
	}

	return &user_service.ListByIdsResponse{
		Total: int32(len(users)),
		Users: UserInfos,
	}, nil
}
