package favorite

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

const (
	FavoriteUnLike = iota
	FavoriteLike
)
const (
	FavoriteVideo = iota
	FavoriteComment
)

// Favorite 点赞
type Favorite struct {
	gorm.Model
	TargetID uint `gorm:"index:idx_target_user_type;not null"`
	UserID   uint `gorm:"index:idx_target_user_type;not null"`
	Type     int  `gorm:"index:idx_target_user_type;not null"` // 0:视频 1:评论
	Status   bool `gorm:"index;not null"`                      // false:未点赞 true:点赞
}

type FavoriteModel struct {
	db *gorm.DB
}

func NewFavoriteModel(db *gorm.DB) *FavoriteModel {
	return &FavoriteModel{db}
}

func (m *FavoriteModel) IsFavoriteVideo(ctx context.Context, videoId, userId uint) (bool, error) {
	return m.isFavorite(ctx, videoId, userId, FavoriteVideo)
}

func (m *FavoriteModel) IsFavoriteComment(ctx context.Context, commentId, userId uint) (bool, error) {
	return m.isFavorite(ctx, commentId, userId, FavoriteComment)
}

func (m *FavoriteModel) isFavorite(ctx context.Context, targetId, userId uint, _type int) (bool, error) {
	var result Favorite
	err := m.db.WithContext(ctx).
		Where("target_id = ? AND user_id = ? AND type = ?", targetId, userId, _type).
		First(&result).Error
	if err != nil {
		return false, err
	}
	return result.Status, nil
}

func (m *FavoriteModel) FavoriteVideo(ctx context.Context, userId, videoId uint, action bool) error {
	return m.favorite(ctx, userId, videoId, FavoriteVideo, action)
}

func (m *FavoriteModel) FavoriteComment(ctx context.Context, userId, commentId uint, action bool) error {
	return m.favorite(ctx, userId, commentId, FavoriteComment, action)
}

func (m *FavoriteModel) favorite(ctx context.Context, userId, targetId uint, _type int, action bool) error {
	result := Favorite{
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
