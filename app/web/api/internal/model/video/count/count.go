package video_count

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

// VideoCountModel 视频播放量统计
type VideoCountModel struct {
	BizRedis *redis.Redis
}

func NewVideoCountModel(rds *redis.Redis) *VideoCountModel {
	return &VideoCountModel{BizRedis: rds}
}

// videoPlayCountKey 视频播放量
func videoPlayCountKey() string {
	return "biz#video#view"
}

// videoShareCountKey 视频分享量
func videoShareCountKey() string {
	return "biz#video#share"
}

// VideoCommentCountKey 视频评论量
func videoCommentCountKey() string {
	return "biz#video#comment"
}

// videoHotVideoKey 热门视频(播放量+分享量)
func videoHotVideoKey() string {
	return "biz#video#hot"
}

// IncrPlayCount 增加播放量
func (m *VideoCountModel) IncrPlayCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZincrbyCtx(ctx, videoPlayCountKey(), 1, strconv.FormatInt(videoId, 10))
	if err != nil {
		return 0, err
	}
	return score, nil
}

// IncrShareCount 增加分享量
func (m *VideoCountModel) IncrShareCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZincrbyCtx(ctx, videoShareCountKey(), 1, strconv.FormatInt(videoId, 10))
	if err != nil {
		return 0, err
	}
	return score, nil
}

// IncrHotCount 增加热度
func (m *VideoCountModel) IncrHotCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZincrbyCtx(ctx, videoHotVideoKey(), 1, strconv.FormatInt(videoId, 10))
	if err != nil {
		return 0, err
	}
	return score, nil
}

// IncrCommentCount 增加评论量
func (m *VideoCountModel) IncrCommentCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZincrbyCtx(ctx, videoCommentCountKey(), 1, strconv.FormatInt(videoId, 10))
	if err != nil {
		return 0, err
	}
	return score, nil
}

// DecrCommentCount 减少评论量
func (m *VideoCountModel) DecrCommentCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZincrbyCtx(ctx, videoCommentCountKey(), -1, strconv.FormatInt(videoId, 10))
	if err != nil {
		return 0, err
	}
	return score, nil
}

// GetCommentCount 获取评论量
func (m *VideoCountModel) GetCommentCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZscoreCtx(ctx, videoCommentCountKey(), strconv.FormatInt(videoId, 10))
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return score, nil
}

func (m *VideoCountModel) GetPlayCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZscoreCtx(ctx, videoPlayCountKey(), strconv.FormatInt(videoId, 10))
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return score, nil
}

func (m *VideoCountModel) GetShareCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZscoreCtx(ctx, videoShareCountKey(), strconv.FormatInt(videoId, 10))
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return score, nil
}

func (m *VideoCountModel) GetHotCount(ctx context.Context, videoId int64) (int64, error) {
	score, err := m.BizRedis.ZscoreCtx(ctx, videoHotVideoKey(), strconv.FormatInt(videoId, 10))
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return score, nil
}

func (m *VideoCountModel) GetHotVideo(ctx context.Context, cursor, pageSize int64) ([]int64, error) {
	limitCtx, err := m.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, videoHotVideoKey(), 0, cursor, 0, int(pageSize))
	if err == redis.Nil {
		return []int64{}, nil
	}
	if err != nil {
		return nil, err
	}
	var videoIds []int64
	for _, pair := range limitCtx {
		videoId, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			continue
		}
		videoIds = append(videoIds, videoId)
	}
	return videoIds, nil
}
