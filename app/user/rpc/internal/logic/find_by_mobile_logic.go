package logic

import (
	"context"
	"errors"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/oss"
	"gorm.io/gorm"

	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/pb/user_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *user_service.FindByMobileRequest) (*user_service.UserInfoResponse, error) {
	user, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Mobile)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 数据库查询失败
		return nil, err
	}
	if user == nil { // 用户不存在
		return &user_service.UserInfoResponse{}, nil
	}
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
	return &user_service.UserInfoResponse{
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
	}, nil
}
