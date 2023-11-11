package video_count

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
)

// VideoCountModel 视频播放量统计
type VideoCountModel struct {
	BizRedis *redis.Redis
}

func NewVideoCountModel(rds *redis.Redis) *VideoCountModel {
	return &VideoCountModel{BizRedis: rds}
}

// videoPlayCountKey 视频播放量
func videoPlayCountKey(videoId int64) string {
	return fmt.Sprintf("biz#video#view#play#%d", videoId)
}

// videoShareCountKey 视频分享量
func videoShareCountKey(videoId int64) string {
	return fmt.Sprintf("biz#video#share#play#%d", videoId)
}

func (m *VideoCountModel) incrementVideoCount(ctx context.Context, key string) (int64, error) {
	count, err := m.BizRedis.IncrCtx(ctx, key)
	if err != nil {
		return 1, err
	}
	return count, nil
}

// IncrVideoPlayCount 视频播放量+1
func (m *VideoCountModel) IncrVideoPlayCount(ctx context.Context, videoId int64) (int64, error) {
	return m.incrementVideoCount(ctx, videoPlayCountKey(videoId))
}

// IncrVideoShareCount 视频分享量+1
func (m *VideoCountModel) IncrVideoShareCount(ctx context.Context, videoId int64) (int64, error) {
	return m.incrementVideoCount(ctx, videoShareCountKey(videoId))
}

// GetVideoPlayCount 获取视频播放量
func (m *VideoCountModel) GetVideoPlayCount(ctx context.Context, videoId int64) (int64, error) {
	return m.getVideoCount(ctx, videoPlayCountKey(videoId))
}

// GetVideoShareCount 获取视频分享量
func (m *VideoCountModel) GetVideoShareCount(ctx context.Context, videoId int64) (int64, error) {
	return m.getVideoCount(ctx, videoShareCountKey(videoId))
}

// GetVideoPlayCount 获取视频播放量
func (m *VideoCountModel) getVideoCount(ctx context.Context, key string) (int64, error) {
	count, err := m.BizRedis.GetCtx(ctx, key)
	if count == "" || err != nil {
		return 0, err
	}
	cnt, err := strconv.ParseInt(count, 10, 64)
	if err != nil {
		return 0, err
	}
	return cnt, nil
}
