package user

import (
	"context"
	user_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailInfoLogic {
	return &GetUserDetailInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailInfoLogic) GetUserDetailInfo(req *types.UserDetailInfoReq) (resp *types.UserDetailInfoResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	contextUserId := req.ContextUserId
	userInfo, err := l.svcCtx.UserRpc.FindById(l.ctx, &user_rpc.FindByIdRequest{
		UserId: contextUserId,
	})

	if err != nil {
		return nil, err
	}
	UserBasicInfo := types.UserBasicInfo{
		ID:            strconv.FormatInt(userInfo.Id, 10),
		NickName:      userInfo.NickName,
		AvatarUrl:     userInfo.AvatarUrl,
		Slogan:        userInfo.Slogan,
		Gender:        int64(userInfo.Gender),
		Age:           int(userInfo.Age),
		BackgroundUrl: userInfo.BackgroundUrl,
	}
	UserStatisticalInfo := types.UserStatisticalInfo{
		FollowingCount:        int(userInfo.FollowingCount),
		FollowerCount:         int(userInfo.FollowerCount),
		LikeCount:             0,
		PublishedVideoCount:   0,
		LikeVideoCount:        0,
		CollectionsVideoCount: 0,
	}
	isFollow := false
	if doerId != 0 {
		//	TODO:判断是否关注
	}
	UserInteractionInfo := types.UserInteractionInfo{
		IsFollow: isFollow,
	}
	resp = &types.UserDetailInfoResp{
		UserBasicInfo:       UserBasicInfo,
		UserInteractionInfo: UserInteractionInfo,
		UserStatisticalInfo: UserStatisticalInfo,
	}
	return
}
