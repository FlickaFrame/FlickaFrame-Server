package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/common"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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
	videoBasicInfo := &types.VideoBasicInfo{}
	err := copier.Copy(videoBasicInfo, video)
	if err != nil {
		logx.Info("copy video to videoBasicInfo fail: ", err)
		return nil, err
	}
	videoBasicInfo.PlayUrl = common.NewURLLogic(c.ctx, c.svcCtx).GetAccessUrl(ctx, video.PlayUrl) // 链接转换
	videoBasicInfo.ThumbUrl = common.NewURLLogic(c.ctx, c.svcCtx).GetAccessUrl(ctx, video.ThumbUrl)
	err = video.LoadAttributes(ctx, c.svcCtx.DB)
	if err != nil {
		logx.Info("loading video attributes from db fail: ", err)
		return nil, err
	}
	// 加载视频作者
	videoBasicInfo.VideoUserInfo = &types.VideoUserInfo{}
	err = copier.Copy(&videoBasicInfo.VideoUserInfo, video.Author)
	videoBasicInfo.VideoUserInfo.ID = strconv.FormatInt(video.AuthorID, 10)
	videoBasicInfo.VideoUserInfo.AvatarUrl = common.NewURLLogic(c.ctx, c.svcCtx).GetAccessUrl(ctx, videoBasicInfo.VideoUserInfo.AvatarUrl)
	if err != nil {
		logx.Info("loading video user fail: ", err)
		return nil, err
	}
	// 加载视频分类
	videoBasicInfo.Category = &types.Category{}
	err = copier.Copy(&videoBasicInfo.Category, video.Category)
	videoBasicInfo.Category.ID = strconv.FormatInt(video.CategoryID, 10)
	if err != nil {
		logx.Info("loading video category fail: ", err)
		return nil, err
	}
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
