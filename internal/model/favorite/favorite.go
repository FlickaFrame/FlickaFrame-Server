package favorite

import (
	"context"
	"errors"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/snowflake"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
	"time"
)

const (
	UnLike = iota
	Like
)
const (
	Video = iota
	Comment
)

// Favorite 点赞
type Favorite struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	TargetID int64 `gorm:"index:idx_target_user_type;not null"`
	UserID   int64 `gorm:"index:idx_target_user_type;not null"`
	Type     int   `gorm:"index:idx_target_user_type;not null"` // 0:视频 1:评论
	Status   bool  `gorm:"index;not null"`                      // false:未点赞 true:点赞
}

type FavoriteModel struct {
	db *orm.DB
}

func NewFavoriteModel(db *orm.DB) *FavoriteModel {
	return &FavoriteModel{db}
}

func (m *FavoriteModel) IsFavoriteVideo(ctx context.Context, videoId, userId int64) (bool, error) {
	return m.isFavorite(ctx, videoId, userId, Video)
}

func (m *FavoriteModel) IsFavoriteComment(ctx context.Context, commentId, userId int64) (bool, error) {
	return m.isFavorite(ctx, commentId, userId, Comment)
}

func (m *FavoriteModel) isFavorite(ctx context.Context, targetId, userId int64, _type int) (bool, error) {
	var result Favorite
	err := m.db.WithContext(ctx).
		Where("target_id = ? AND user_id = ? AND type = ?", targetId, userId, _type).
		First(&result).Error
	if err != nil {
		return false, err
	}
	return result.Status, nil
}

func (m *FavoriteModel) FavoriteVideo(ctx context.Context, userId, videoId int64, action bool) error {
	return m.favorite(ctx, userId, videoId, Video, action)
}

func (m *FavoriteModel) FavoriteComment(ctx context.Context, userId, commentId int64, action bool) error {
	return m.favorite(ctx, userId, commentId, Comment, action)
}

func (m *FavoriteModel) favorite(ctx context.Context, userId, targetId int64, _type int, action bool) error {
	result := Favorite{
		ID:       snowflake.FavoriteIDNode.Generate().Int64(),
		TargetID: targetId,
		UserID:   userId,
		Type:     _type,
		Status:   action,
	}
	// 尝试插入
	err := m.db.WithContext(ctx).
		Where("target_id ? AND user_id = ? AND type = ?", targetId, userId, _type).
		First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) { // 未找到记录
		err = m.db.WithContext(ctx).Create(&result).Error
		if err != nil {
			return err
		}
		return nil
	} else { // 找到记录
		err = m.db.WithContext(ctx).
			Where("target_id ? AND user_id = ? AND type = ?", targetId, userId, _type).
			Update("status", action).Error
		if err != nil {
			return err
		}
		return nil
	}
}
