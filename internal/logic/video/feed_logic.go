package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/jinzhu/copier"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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
	if req.Cursor != 0 {
		LatestTime = time.UnixMilli(req.Cursor)
	}
	videos, err := l.svcCtx.VideoModel.List(l.ctx, video_model.ListOption{
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
	videoRsp := make([]*types.VideoItem, 0, len(videos))
	for _, v := range videos { // TODO: 优化
		feedItem := &types.VideoItem{}
		author := l.svcCtx.UserModel.MustFindOne(l.ctx, v.AuthorID)
		copier.Copy(feedItem, v)
		feedItem.PublishTime = v.CreatedAt.Format("2006-01-02 15:04:05")
		feedItem.VideoUrl = storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, v.PlayUrl)
		feedItem.User.Avatar = storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, author.AvatarUrl)
		videoRsp = append(videoRsp, feedItem)
		feedItem.Interaction.IsFollow = l.svcCtx.UserModel.IsFollowing(l.ctx, doerId, v.AuthorID)
		feedItem.Interaction.Liked, _ = l.svcCtx.FavoriteModel.IsFavoriteVideo(l.ctx, doerId, v.ID)
	}

	// 判断是否无视频
	nextTime := LatestTime.UnixMilli()
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreatedAt.UnixMilli() - 1
	}
	resp = &types.FeedResp{
		List: videoRsp,
		Next: strconv.FormatInt(nextTime, 10),
	}
	return
}
