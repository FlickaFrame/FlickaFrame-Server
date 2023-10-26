package feed

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/jinzhu/copier"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	LatestTime := time.Now()
	if req.LatestTime != 0 {
		LatestTime = time.UnixMilli(req.LatestTime)
	}
	videos, err := l.svcCtx.VideoModel.FindByLatestTime(l.ctx, LatestTime, 10)
	if err != nil {
		l.Logger.Infof("find videos by latest time error: %v", err)
		return
	}
	videoRsp := make([]*types.Video, 0, len(videos))
	for _, v := range videos { // TODO: 优化
		feedItem := &types.Video{}
		author := l.svcCtx.UserModel.MustFindOne(l.ctx, v.AuthorID)
		copier.Copy(feedItem, v)
		copier.Copy(&feedItem.Author, author)
		feedItem.PlayUrl = l.GetVideoURL(v.PlayUrl)
		feedItem.Author.AvatarUrl = l.GetVideoURL(author.AvatarUrl)
		videoRsp = append(videoRsp, feedItem)
	}

	// 判断是否无视频
	nextTime := LatestTime.UnixMilli()
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreatedAt.UnixMilli() - 1
	}
	resp = &types.FeedResp{
		VideoList: videoRsp,
		NextTime:  nextTime,
		Length:    len(videos),
	}
	return
}

func (l *FeedLogic) GetVideoURL(key string) string {
	privateAccessURL := storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, key)
	return privateAccessURL
}
