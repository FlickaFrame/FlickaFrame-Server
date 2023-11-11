package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HotLogic {
	return &HotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HotLogic) Hot(req *types.FeedReq) (resp *types.FeedResp, err error) {
	videoConvert := NewConvert(l.ctx, l.svcCtx)
	return videoConvert.Feed(req, HotFeed)
}

func HotFeed(ctx context.Context, svcCtx *svc.ServiceContext, req *types.FeedReq) ([]*video_model.Video, error) {
	videoCount := video_count.NewVideoCountModel(svcCtx.BizRedis)
	videoIds, err := videoCount.GetHotVideo(ctx, req.Cursor, int64(req.Limit))
	if err != nil {
		return nil, err
	}
	return svcCtx.VideoModel.FindByIDs(ctx, videoIds)
}
