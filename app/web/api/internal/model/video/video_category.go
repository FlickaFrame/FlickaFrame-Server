package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/base"
)

type Category struct {
	base.Model

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
