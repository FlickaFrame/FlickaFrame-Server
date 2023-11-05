package video

import (
	"context"
	"time"
)

type Category struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name   string   `gorm:"unique"`
	Videos []*Video `gorm:"-"`
}

func (c *Category) TableName() string {
	return "video_category"
}

func (m *VideoModel) FindCategories(ctx context.Context) ([]*Category, error) {
	var result []*Category
	err := m.db.WithContext(ctx).Find(&result).Error
	return result, err
}
