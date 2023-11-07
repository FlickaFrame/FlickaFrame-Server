package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/base"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
)

type FavoriteType = string

const (
	VideoFavoriteType   = "video"
	CommentFavoriteType = "comment"
)

// Favorite 点赞
type Favorite struct {
	base.Model
	TargetID int64 `gorm:"index:idx_target_user;not null"`
	UserID   int64 `gorm:"index:idx_target_user;not null"`
	Type     string
}

type Model struct {
	db *orm.DB
}

func NewFavoriteModel(db *orm.DB) *Model {
	return &Model{db}
}

func (m *Model) IsExist(ctx context.Context, targetId, userId int64) error {
	return m.db.WithContext(ctx).
		Where("target_id = ? AND user_id = ?", targetId, userId).
		First(&Favorite{}).Error
}

func (m *Model) Delete(ctx context.Context, userId, targetId int64) error {
	return m.db.WithContext(ctx).
		Where("target_id = ? AND user_id = ?", targetId, userId).
		Delete(&Favorite{}).Error
}

func (m *Model) Create(ctx context.Context, userId, targetId int64, typ string) error {
	result := Favorite{
		Model:    base.NewModel(),
		TargetID: targetId,
		UserID:   userId,
		Type:     typ,
	}
	return m.db.WithContext(ctx).Create(&result).Error
}
