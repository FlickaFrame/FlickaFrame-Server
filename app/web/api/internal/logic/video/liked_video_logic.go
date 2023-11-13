package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type LikedVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikedVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikedVideoLogic {
	return &LikedVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikedVideoLogic) LikedVideo(req *types.FeedReq) (resp *types.FeedResp, err error) {
	videoConvert := NewConvert(l.ctx, l.svcCtx)
	return videoConvert.Feed(req, LikedVideo)
}

func LikedVideo(ctx context.Context, svcCtx *svc.ServiceContext, req *types.FeedReq) ([]*video_model.Video, error) {
	doerId := jwt.GetUidFromCtx(ctx)
	videoIds, err := svcCtx.FavoriteModel.FindVideoIdsByUser(ctx, doerId, req.Cursor, req.Limit)
	if err != nil {
		logx.WithContext(ctx).Errorf("FindVideoIdsByUser: find videos by latest time error: %v", err)
		return nil, err
	}
	// 获取视频基本信息
	return svcCtx.VideoModel.FindByIDs(ctx, videoIds)
}
