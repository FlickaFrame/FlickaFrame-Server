package comment

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string            `gorm:"type:varchar(255);not null"`
	UserID  int               `gorm:"index:idx_userid;not null"`
	User    user_model.User   `gorm:"-"`
	VideoID int               `gorm:"index:idx_videoid;not null"`
	Video   video_model.Video `gorm:"-"`
}

func (c *Comment) TableName() string {
	return "comment"
}

type CommentModel struct {
	db *gorm.DB
}

func (c *CommentModel) Insert(ctx context.Context, comment Comment) error {
	return c.db.WithContext(ctx).Create(&comment).Error
}

func (c *CommentModel) Delete(ctx context.Context, comment Comment) error {
	return c.db.WithContext(ctx).Delete(&comment).Error
}

func (c *CommentModel) Update(ctx context.Context, comment Comment) error {
	return c.db.WithContext(ctx).Save(&comment).Error
}
