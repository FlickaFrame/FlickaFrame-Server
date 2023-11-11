package video

import (
	"context"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoInfoLogic {
	return &GetVideoInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoInfoLogic) GetVideoInfo(req *types.GetVideoInfoReq) (resp *types.GetVideoInfoResp, err error) {
	video, err := l.svcCtx.VideoModel.FindOne(l.ctx, req.VideoId)
	if err != nil {
		logx.Info(err)
		return nil, err
	}
	videoBasicInfo, err := NewConvert(l.ctx, l.svcCtx).BuildVideoBasicInfo(l.ctx, video)
	resp = &types.GetVideoInfoResp{
		Video: &types.VideoInfoItem{
			VideoBasicInfo:       *videoBasicInfo,
			VideoStatisticalInfo: types.VideoStatisticalInfo{},
			VideoInteractInfo:    types.VideoInteractInfo{},
		},
	}
	// 互动信息(与当前登录用户的关注关系)
	doerId := jwt.GetUidFromCtx(l.ctx)
	resp.Video.IsFavorite, err = l.svcCtx.FavoriteModel.IsExist(l.ctx, video.ID, doerId)
	follow, err := l.svcCtx.FollowRpc.IsFollow(l.ctx, &follow_rpc.IsFollowReq{
		UserId:         doerId,
		FollowedUserId: video.AuthorID,
	})
	if err != nil {
		return nil, err
	}
	resp.Video.VideoUserInfo.IsFollow = follow.IsFollow
	return
}
