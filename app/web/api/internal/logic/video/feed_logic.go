package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
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
	var (
		doerId       = jwt.GetUidFromCtx(l.ctx) // 从context中获取当前用户id
		videoConvert = NewConvert(l.ctx, l.svcCtx)
		emptyResp    = &types.FeedResp{ // 默认空列表
			List:  make([]*types.FeedVideoItem, 0),
			IsEnd: true,
			Next:  strconv.FormatInt(time.Now().UnixMilli(), 10),
		}
	)
	if req.Cursor == 0 {
		req.Cursor = time.Now().UnixMilli()
	}
	videos, err := l.svcCtx.VideoModel.List(l.ctx, video_model.ListOption{
		AuthorID:   util.MustString2Int64(req.AuthorID),
		LatestTime: time.UnixMilli(req.Cursor),
		Limit:      req.Limit,
		QueryAll:   false,
		CategoryID: util.MustString2Int64(req.CategoryID),
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("VideoModel.List : find videos by latest time error: %v", err)
		return emptyResp, err
	}
	if len(videos) == 0 {
		return emptyResp, nil
	}

	feedVideos := make([]*types.FeedVideoItem, 0, len(videos))
	for i, video := range videos {
		item := &types.FeedVideoItem{
			VideoBasicInfo: types.VideoBasicInfo{
				ID:            strconv.FormatInt(video.ID, 10),
				Title:         video.Title,
				Description:   video.Description,
				CreatedAt:     video.CreatedAt.UnixMilli(),
				VideoHeight:   video.VideoHeight,
				VideoWidth:    video.VideoWidth,
				VideoDuration: video.VideoDuration,
			},
			VideoStatisticalInfo: types.VideoStatisticalInfo{},
			VideoInteractInfo: types.VideoInteractInfo{
				IsFavorite: false,
			},
		}
		// 加载播放链接
		item.PlayUrl = videoConvert.WithVideoURL(l.ctx, video.PlayUrl)
		// 加载封面链接
		item.ThumbUrl = videoConvert.WithVideoURL(l.ctx, video.ThumbUrl)
		// 加载视频作者
		author, err := videoConvert.WithVideoAuthor(l.ctx, doerId, videos[i].AuthorID)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("VideoModel.WithVideoAuthor : find video author by id error: %v", err)
			return nil, err
		}
		item.VideoUserInfo = author
		// 加载视频分类
		category, err := videoConvert.WithVideoCategory(l.ctx, videos[i].CategoryID)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("VideoModel.WithVideoCategory : find video category by id error: %v", err)
			return nil, err
		}
		item.Category = category
		// 加载视频标签
		tags, err := l.svcCtx.VideoModel.FindTagsByVideoId(l.ctx, videos[i].ID)
		for _, tag := range tags {
			item.Tags = append(item.Tags, &types.Tag{
				Id:   strconv.FormatInt(tag.ID, 10),
				Name: tag.Name,
			})
		}
		if item.Tags == nil {
			// 防止前端解析出错
			item.Tags = make([]*types.Tag, 0)
		}
		if err != nil {
			logx.WithContext(l.ctx).Errorf("VideoModel.FindTagsByVideoId : find video tags by video id error: %v", err)
			return nil, err
		}
		// 加载视频点赞数
		favoriteCount, err := l.svcCtx.FavoriteModel.CountByVideoId(l.ctx, videos[i].ID)
		if err != nil {
			return nil, err
		}
		item.FavoriteCount = favoriteCount
		// 加载视频评论数
		commentCount, err := l.svcCtx.CommentModel.CountByVideoId(l.ctx, videos[i].ID)
		if err != nil {
			return nil, err
		}
		item.CommentCount = commentCount
		// 加载视频分享数
		shareCount, err := videoConvert.WithVideoShareCount(l.ctx, videos[i].ID)
		if err != nil {
			return nil, err
		}
		item.ShareCount = shareCount
		// 加载视频播放数
		viewCount, err := videoConvert.WithVideoPlayCount(l.ctx, videos[i].ID)
		if err != nil {
			return nil, err
		}
		item.ViewCount = viewCount
		// 加载视频是否点赞
		exist, err := videoConvert.WithVideoIsFavorite(l.ctx, videos[i].ID, doerId)
		if err != nil {
			return nil, err
		}
		item.IsFavorite = exist
		feedVideos = append(feedVideos, item)
	}
	resp = &types.FeedResp{
		Next:  strconv.FormatInt(videos[len(videos)-1].CreatedAt.UnixMilli(), 10),
		List:  feedVideos,
		IsEnd: len(feedVideos) < req.Limit,
	}
	return resp, nil
}
