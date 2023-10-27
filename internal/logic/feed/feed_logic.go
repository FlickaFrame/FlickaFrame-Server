package feed

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
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
	doerId := jwt.GetUidFromCtx(l.ctx) // 从context中获取当前用户id
	LatestTime := time.Now()
	if req.LatestTime != 0 {
		LatestTime = time.UnixMilli(req.LatestTime)
	}
	videos, err := l.svcCtx.VideoModel.List(l.ctx, video_model.Option{
		AuthorID:   req.AuthorID,
		LatestTime: LatestTime,
		Limit:      req.Limit,
		QueryAll:   false,
		CategoryID: req.CategoryID,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("find videos by latest time error: %v", err)
		return
	}
	videoRsp := make([]*types.Video, 0, len(videos))
	for _, v := range videos { // TODO: 优化
		feedItem := &types.Video{}
		author := l.svcCtx.UserModel.MustFindOne(l.ctx, v.AuthorID)
		copier.Copy(feedItem, v)
		feedItem.CreatedAt = v.CreatedAt.Format("2006-01-02 15:04:05")
		copier.Copy(&feedItem.Author, author)
		feedItem.PlayUrl = storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, v.PlayUrl)
		feedItem.Author.AvatarUrl = storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, author.AvatarUrl)
		videoRsp = append(videoRsp, feedItem)
		feedItem.IsFollow, _ = l.svcCtx.FollowModel.IsFollow(l.ctx, doerId, v.AuthorID)
		feedItem.IsFav, _ = l.svcCtx.FavoriteModel.IsFavorite(l.ctx, doerId, v.ID)
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
