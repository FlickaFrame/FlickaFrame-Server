package user

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/jinzhu/copier"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type Convert struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvert(ctx context.Context, svcCtx *svc.ServiceContext) *Convert {
	return &Convert{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *Convert) BuildUserDetailInfo(ctx context.Context, doer *user_model.User, user *user_model.User) (*types.UserDetailInfoResp, error) {
	var userInfo types.UserDetailInfoResp
	_ = copier.Copy(&userInfo, user)
	userInfo.AvatarUrl = c.getAccessUrl(ctx, user.AvatarUrl)
	if doer != user {
		// TODO: 做权限判断用户属性可见性&交互信息获取
		//userInfo.IsFollow = c.svcCtx.UserModel.IsFollowing(ctx, doer.ID, user.ID)
	}
	// TODO: 通过redis缓存获取统计信息
	//userInfo.UserStatisticalInfo = types.UserStatisticalInfo{
	//	FollowingCount:        0, //关注数
	//	FollowerCount:         0, //粉丝数

	//	LikeCount:             0, //获赞数
	//	PublishedVideoCount:   0, //作品数
	//	LikeVideoCount:        0, //点赞数
	//	CollectionsVideoCount: 0, //收藏数
	//}
	return &userInfo, nil
}

func (c *Convert) BuildUserBasicInfo(ctx context.Context, user *user_model.User) (*types.UserBasicInfo, error) {
	var userInfo types.UserBasicInfo
	err := copier.Copy(&userInfo, user)
	if err != nil {
		logx.Info("copy user to userInfo fail: ", err)
	}
	userInfo.ID = strconv.FormatInt(user.ID, 10)
	userInfo.AvatarUrl = c.getAccessUrl(ctx, user.AvatarUrl)
	return &userInfo, nil
}

func (c *Convert) BuildUserInteractionInfo(ctx context.Context, doer, contextUser int64) (*types.UserInteractionInfo, error) {
	// 关注信息
	isFollow := c.svcCtx.UserModel.IsFollowing(ctx, doer, contextUser)
	return &types.UserInteractionInfo{
		IsFollow: isFollow,
	}, nil
}

func (c *Convert) BuildUserInteractionInfoList(ctx context.Context, doer int64, userList []*user_model.User) ([]*types.UserInteractionInfo, error) {
	var userInfoList []*types.UserInteractionInfo
	for _, user := range userList {
		userInfo, err := c.BuildUserInteractionInfo(ctx, doer, user.ID)
		if err != nil {
			logx.Info("build user interaction info fail: ", err)
			return nil, err
		}
		userInfoList = append(userInfoList, userInfo)
	}
	return userInfoList, nil
}

func (c *Convert) BuildUserStatisticalInfo(ctx context.Context, user *user_model.User) (*types.UserStatisticalInfo, error) {
	return &types.UserStatisticalInfo{
		FollowingCount:        0,
		FollowerCount:         0,
		LikeCount:             0,
		PublishedVideoCount:   0,
		LikeVideoCount:        0,
		CollectionsVideoCount: 0,
	}, nil
}

func (c *Convert) buildUserBasicInfoList(ctx context.Context, userList []*user_model.User) ([]*types.UserBasicInfo, error) {
	var userInfoList []*types.UserBasicInfo
	for _, user := range userList {
		userInfo, err := c.BuildUserBasicInfo(ctx, user)
		if err != nil {
			logx.Info("build user basic info fail: ", err)
			return nil, err
		}
		userInfoList = append(userInfoList, userInfo)
	}
	return userInfoList, nil
}

func (c *Convert) getAccessUrl(ctx context.Context, key string) string {
	return storage.MakePublicURL(c.svcCtx.Config.Oss.Endpoint, key)
}
