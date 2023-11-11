package video_history

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const CacheMaxPlayHistoryCount = 1000

type PlayHistory struct {
	VideoId  int64     `json:"videoId"`  // 视频ID
	PlayTime time.Time `json:"playTime"` // 播放时间
}

type VideoHistoryModel struct {
	BizRedis *redis.Redis
}

func NewVideoHistoryModel(rds *redis.Redis) *VideoHistoryModel {
	return &VideoHistoryModel{BizRedis: rds}
}

func userPlayHistoryKey(userId int64) string {
	return fmt.Sprintf("biz#user#paly_history#%d", userId)
}

// GetPlayHistoryFromCache retrieves play history for a user from Redis ZSET
func (m *VideoHistoryModel) GetPlayHistoryFromCache(ctx context.Context, userId, cursor, pageSize int64) ([]*PlayHistory, error) {
	key := userPlayHistoryKey(userId)
	exist, err := m.BizRedis.ExistsCtx(ctx, key)
	if err != nil || !exist {
		logx.Errorf("[GetPlayHistoryFromCache] BizRedis.ExistsCtx error: %v", err)
		return nil, err
	}
	if !exist {
		return nil, nil
	}
	pairs, err := m.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, key, 0, cursor, 0, int(pageSize))
	if err != nil {
		logx.Errorf("[GetPlayHistoryFromCache] BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx error: %v", err)
		return nil, err
	}
	var playHistories []*PlayHistory
	for _, pair := range pairs {
		VideoId, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			logx.Errorf("[GetPlayHistoryFromCache] strconv.ParseInt error: %v", err)
			continue
		}
		playHistories = append(playHistories, &PlayHistory{
			VideoId:  VideoId,
			PlayTime: time.UnixMilli(pair.Score),
		})
	}
	return playHistories, nil
}
func (m *VideoHistoryModel) AddCachePlayHistory(ctx context.Context, userId int64, videoId int64) error {
	var histories []*PlayHistory
	histories = append(histories,
		&PlayHistory{
			VideoId:  videoId,
			PlayTime: time.Now(),
		})
	return m.AddCachePlayHistories(ctx, userId, histories)
}

// AddCachePlayHistories 将播放记录写入缓存
func (m *VideoHistoryModel) AddCachePlayHistories(ctx context.Context, userId int64, histories []*PlayHistory) error {
	if len(histories) == 0 {
		return nil
	}
	key := userPlayHistoryKey(userId)
	exist, err := m.BizRedis.ExistsCtx(ctx, key)
	if err != nil {
		logx.Errorf("[AddCachePlayHistories] BizRedis.ExistsCtx error: %v", err)
		return err
	}
	if !exist { // 添加结束标志
		histories = append(histories, &PlayHistory{VideoId: -1})
	}
	for _, history := range histories { // 写入redis中
		var score int64
		if history.VideoId == -1 {
			score = 0
		} else {
			score = history.PlayTime.UnixMilli()
		}
		_, err = m.BizRedis.ZaddCtx(ctx, key, score, strconv.FormatInt(history.VideoId, 10))
		if err != nil {
			logx.Errorf("[addCachePlayHistories] BizRedis.ZaddCtx error: %v", err)
			return err
		}
	}
	// 使用ZREMRANGEBYRANK命令剔除多余的关注关系
	_, err = m.BizRedis.ZremrangebyrankCtx(ctx, key, 0, -(CacheMaxPlayHistoryCount + 1))
	if err != nil {
		logx.Errorf("[AddCachePlayHistories] Redis Zremrangebyrank error: %v", err)
	}
	return nil
}

func (m *VideoHistoryModel) ClearPlayHistory(ctx context.Context, userId int64) error {
	key := userPlayHistoryKey(userId)
	_, err := m.BizRedis.DelCtx(ctx, key)
	return err
}

func (m *VideoHistoryModel) DeleteSpecificPlayHistory(ctx context.Context, userId, videoId int64) error {
	key := userPlayHistoryKey(userId)
	member := fmt.Sprintf("%d", videoId)
	_, err := m.BizRedis.ZremCtx(ctx, key, member)
	return err
}
