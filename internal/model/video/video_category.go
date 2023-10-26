package video

import (
	"context"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
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

func (m *VideoModel) FindVideoByCategory(ctx context.Context, categoryID int64, limit int) ([]*Video, error) {
	var result []*Video
	err := m.db.WithContext(ctx).
		Where("category_id = ?", categoryID).
		Order("id desc").
		Limit(limit).
		Find(&result).Error
	return result, err
}
