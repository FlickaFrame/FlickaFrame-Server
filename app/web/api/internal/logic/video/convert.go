package video

import (
	"context"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/oss"
	user_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/logic/user"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type Convert struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvert(ctx context.Context, svcCtx *svc.ServiceContext) *Convert {
	return &Convert{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BuildVideoBasicInfo 用于构建视频基本信息
func (c *Convert) BuildVideoBasicInfo(ctx context.Context, video *video_model.Video) (*types.VideoBasicInfo, error) {
	var (
		videoBasicInfo = &types.VideoBasicInfo{}
		userConvert    = user.NewConvert(c.ctx, c.svcCtx)
	)
	err := copier.Copy(videoBasicInfo, video)
	if err != nil {
		logx.Info("copy video to videoBasicInfo fail: ", err)
		return nil, err
	}

	videoBasicInfo.PlayUrl = userConvert.GetAccessUrl(ctx, video.PlayUrl) // 链接转换
	videoBasicInfo.ThumbUrl = userConvert.GetAccessUrl(ctx, video.ThumbUrl)
	videoBasicInfo.CreatedAt = video.CreatedAt.UnixMilli() // 时间转换

	video.Author = c.svcCtx.UserModel.MustFindOne(ctx, video.AuthorID)
	video.Category = c.svcCtx.VideoModel.MustFindOneCategory(ctx, video.CategoryID)
	video.Tags = c.svcCtx.VideoModel.MustFindTagsByVideoId(ctx, video.ID)
	if err != nil {
		logx.Info("loading video attributes from db fail: ", err)
		return nil, err
	}
	// 加载视频作者
	videoBasicInfo.VideoUserInfo = &types.VideoUserInfo{}
	err = copier.Copy(&videoBasicInfo.VideoUserInfo, video.Author)
	if err != nil {
		logx.Info("loading video user fail: ", err)
		return nil, err
	}
	videoBasicInfo.VideoUserInfo.ID = strconv.FormatInt(video.AuthorID, 10)
	videoBasicInfo.VideoUserInfo.AvatarUrl = user.NewConvert(c.ctx, c.svcCtx).GetAccessUrl(ctx, videoBasicInfo.VideoUserInfo.AvatarUrl)

	// 加载视频分类
	videoBasicInfo.Category = &types.Category{}
	err = copier.Copy(&videoBasicInfo.Category, video.Category)
	if err != nil {
		logx.Info("loading video category fail: ", err)
		return nil, err
	}
	videoBasicInfo.Category.ID = strconv.FormatInt(video.CategoryID, 10)

	// 加载视频标签
	videoBasicInfo.Tags = make([]*types.Tag, len(video.Tags))
	err = copier.Copy(&videoBasicInfo.Tags, video.Tags)
	if err != nil {
		logx.Info("loading video tags fail: ", err)
		return nil, err
	}
	return videoBasicInfo, err
}

// BuildVideoBasicInfoList 用于构建视频基本信息列表
func (c *Convert) BuildVideoBasicInfoList(ctx context.Context, videoList []*video_model.Video) ([]*types.VideoBasicInfo, error) {
	var videoInfoList []*types.VideoBasicInfo
	for _, video := range videoList {
		videoInfo, err := c.BuildVideoBasicInfo(ctx, video)
		if err != nil {
			return nil, err
		}
		videoInfoList = append(videoInfoList, videoInfo)
	}
	return videoInfoList, nil
}

// WithVideoIsFavorite 用于加载视频是否被当前用户点赞
func (c *Convert) WithVideoIsFavorite(ctx context.Context, videoId, doerId int64) (bool, error) {
	return c.svcCtx.FavoriteModel.IsExist(ctx, videoId, doerId)
}

// WithVideoPlayCount 用于加载播放量
func (c *Convert) WithVideoPlayCount(ctx context.Context, videoId int64) (int64, error) {
	count, err := video_count.NewVideoCountModel(c.svcCtx.BizRedis).
		GetVideoPlayCount(ctx, videoId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// WithVideoShareCount 用于加载分享量
func (c *Convert) WithVideoShareCount(ctx context.Context, videoId int64) (int64, error) {
	count, err := video_count.NewVideoCountModel(c.svcCtx.BizRedis).
		GetVideoShareCount(ctx, videoId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// WithVideoTags 用于加载视频标签信息
func (c *Convert) WithVideoTags(ctx context.Context, tagIds []int64) ([]*types.Tag, error) {
	var tagsInfo []*types.Tag
	tags, err := c.svcCtx.VideoModel.FindTagsByIds(ctx, tagIds)
	if err != nil {
		return tagsInfo, err
	}
	for _, tag := range tags {
		tagsInfo = append(tagsInfo, &types.Tag{
			Id:   strconv.FormatInt(tag.ID, 10),
			Name: tag.Name,
		})
	}
	return tagsInfo, nil
}

// WithVideoCategory 用于加载视频分类信息
func (c *Convert) WithVideoCategory(ctx context.Context, categoryId int64) (*types.Category, error) {
	category, err := c.svcCtx.VideoModel.FindOneCategory(ctx, categoryId)
	if err != nil {
		logx.Info("FindOneCategory: find category by id fail: ", err)
		return nil, err
	}
	return &types.Category{
		ID:   strconv.FormatInt(category.ID, 10),
		Name: category.Name,
	}, nil
}

// WithVideoLiked 用于加载视频点赞信息(用户是否点赞)
func (c *Convert) WithVideoLiked(ctx context.Context, doer, videoId int64) (bool, error) {
	isLiked, err := c.svcCtx.FavoriteModel.IsExist(ctx, videoId, doer)
	if err != nil {
		logx.Info("IsExist: find favorite by id fail: ", err)
		return false, err
	}
	return isLiked, nil
}

// WithVideoAuthor 用于加载视频作者信息(用户是否关注作者)
func (c *Convert) WithVideoAuthor(ctx context.Context, doer, userId int64) (*types.VideoUserInfo, error) {
	userInfo, err := c.svcCtx.UserRpc.FindById(ctx, &user_rpc.FindByIdRequest{
		UserId: userId,
	})
	// 加载视频作者基本信息
	if err != nil {
		logx.Info("FindById: find user by id fail: ", err)
		return nil, err
	}
	videoUserInfo := &types.VideoUserInfo{
		ID:        strconv.FormatInt(userInfo.Id, 10),
		NickName:  userInfo.NickName,
		AvatarUrl: userInfo.AvatarUrl,
		Slogan:    userInfo.Slogan,
		Gender:    int64(userInfo.Gender),
		IsFollow:  false,
	}
	// 加载视频作者与当前用户的关注关系
	if doer != 0 {
		isFollow, err := c.svcCtx.FollowRpc.IsFollow(ctx, &follow_rpc.IsFollowReq{
			UserId:         doer,
			FollowedUserId: userId,
		})
		if err != nil {
			logx.Info("IsFollow: find follow by id fail: ", err)
			return nil, err
		}
		videoUserInfo.IsFollow = isFollow.IsFollow
	}
	return videoUserInfo, nil
}

// WithVideoURL 用于加载视频播放地址
func (c *Convert) WithVideoURL(ctx context.Context, key string) string {
	url, err := c.svcCtx.OssRpc.GetFileAccessUrl(ctx, &oss.GetFileAccessUrlRequest{
		Key: key,
	})
	if err != nil {
		logx.Info("GetFileAccessUrl: get file access url fail: ", err)
		return url.Url
	}
	return url.Url
}
