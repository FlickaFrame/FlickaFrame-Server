package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentUserInfoLogic {
	return &CurrentUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CurrentUserInfo 获取当前用户信息
func (l *CurrentUserInfoLogic) CurrentUserInfo(req *types.UserDetailInfoReq) (resp *types.UserDetailInfoResp, err error) {
	req.ContextUserId = jwt.GetUidFromCtx(l.ctx)
	doerId := jwt.GetUidFromCtx(l.ctx)
	userInfo, err := l.svcCtx.UserRpc.FindById(l.ctx, &user.FindByIdRequest{
		UserId: doerId,
	})
	return &types.UserDetailInfoResp{
		UserBasicInfo: types.UserBasicInfo{
			ID:            strconv.FormatInt(userInfo.Id, 10),
			NickName:      userInfo.NickName,
			AvatarUrl:     userInfo.AvatarUrl,
			Slogan:        userInfo.Slogan,
			Gender:        int64(userInfo.Gender),
			Age:           int(userInfo.Age),
			BackgroundUrl: userInfo.BackgroundUrl,
		},
		UserStatisticalInfo: types.UserStatisticalInfo{
			FollowingCount:        int(userInfo.FollowingCount),
			FollowerCount:         int(userInfo.FollowerCount),
			LikeCount:             0,
			PublishedVideoCount:   0,
			LikeVideoCount:        0,
			CollectionsVideoCount: 0,
		},
		UserInteractionInfo: types.UserInteractionInfo{
			IsFollow: false,
		},
	}, nil
}
