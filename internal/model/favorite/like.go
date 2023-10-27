package favorite

import (
	"context"
	"gorm.io/gorm"
)

const (
	FavoriteUnLike = iota
	FavoriteLike
)

// Favorite 点赞
type Favorite struct {
	gorm.Model
	TargetID uint `gorm:"index:idx_target_user_type;not null"`
	UserID  uint `gorm:"index:idx_target_user_type;not null"`
	Type    int `gorm:"index:idx_target_user_type;not null"`      // 0:视频 1:评论
	Status  bool `gorm:"index;not null"` // false:未点赞 true:点赞
}

type FavoriteModel struct {
	db *gorm.DB
}

func NewFavoriteModel(db *gorm.DB) *FavoriteModel {
	return &FavoriteModel{db}
}

func (m *FavoriteModel) IsFavorite(ctx context.Context, videoId, userId uint) (bool, error) {
	if videoId == 0 || userId == 0 {
		return false, nil
	}
	var result Favorite
	err := m.db.WithContext(ctx).Where("video_id = ? AND user_id = ?", videoId, userId).First(&result).Error
	if err != nil {
		return false, err
	}
	return result.Status, nil
}

// Insert 新增点赞状态
func (m *FavoriteModel) Insert(ctx context.Context, like *Favorite) error {
	return m.db.WithContext(ctx).Create(like).Error
}

// Update 更新点赞状态
func (m *FavoriteModel) Update(ctx context.Context, status bool) error {
	return m.db.WithContext(ctx).Update("status", status).Error
}
