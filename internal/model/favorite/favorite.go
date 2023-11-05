package favorite

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/snowflake"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"time"
)

// Favorite 点赞
type Favorite struct {
	ID        int64 `gorm:"primary_key"`
	TargetID  int64 `gorm:"index:idx_target_user;not null"`
	UserID    int64 `gorm:"index:idx_target_user;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
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

func (m *Model) Create(ctx context.Context, userId, targetId int64) error {
	result := Favorite{
		ID:       snowflake.FavoriteIDNode.Generate().Int64(),
		TargetID: targetId,
		UserID:   userId,
	}
	return m.db.WithContext(ctx).Create(&result).Error
}
