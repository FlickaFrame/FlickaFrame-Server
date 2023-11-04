package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/common"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type convert struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func newConvert(ctx context.Context, svcCtx *svc.ServiceContext) *convert {
	return &convert{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BuildVideoBasicInfo 用于构建视频基本信息
func (c *convert) BuildVideoBasicInfo(ctx context.Context, video *video_model.Video) (*types.VideoBasicInfo, error) {
	videoBasicInfo := &types.VideoBasicInfo{}
	err := copier.Copy(videoBasicInfo, video)
	// 链接转换
	videoBasicInfo.PlayUrl = common.NewURLLogic(c.ctx, c.svcCtx).GetAccessUrl(ctx, video.PlayUrl)
	videoBasicInfo.ThumbUrl = common.NewURLLogic(c.ctx, c.svcCtx).GetAccessUrl(ctx, video.ThumbUrl)
	// 时间转换
	videoBasicInfo.CreatedAt = video.CreatedAt.UnixMilli()

	err = video.LoadAttributes(ctx, c.svcCtx.DB)
	if err != nil {
		logx.Info("loading video attributes from db fail: ", err)
		return nil, err
	}
	// 加载视频作者
	videoBasicInfo.VideoUserInfo = &types.VideoUserInfo{}
	err = copier.Copy(&videoBasicInfo.VideoUserInfo, video.Author)
	videoBasicInfo.VideoUserInfo.AvatarUrl = common.NewURLLogic(c.ctx, c.svcCtx).GetAccessUrl(ctx, videoBasicInfo.VideoUserInfo.AvatarUrl)
	if err != nil {
		logx.Info("loading video user fail: ", err)
		return nil, err
	}
	// 加载视频分类
	videoBasicInfo.Category = &types.Category{}
	err = copier.Copy(&videoBasicInfo.Category, video.Category)
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
func (c *convert) BuildVideoBasicInfoList(ctx context.Context, videoList []*video_model.Video) ([]*types.VideoBasicInfo, error) {
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
