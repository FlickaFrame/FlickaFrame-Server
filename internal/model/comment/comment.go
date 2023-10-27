package comment

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content       string            `gorm:"not null"`  			               // 评论内容
	OwnerUID      int               `gorm:"index:idx_owneruid;not null"`   // 评论者ID
	TargetID      int               `gorm:"index:idx_targetid;not null"`   // 评论目标ID
	User          user_model.User   `gorm:"-"`
	Video         video_model.Video `gorm:"-"`
	ReplayCount   int               `gorm:"default:0"`                     // 回复数
	LikeCount     int               `gorm:"default:0"`                     // 点赞数
	Type          int               `gorm:"default:0"`                     // 评论类型(0:评论视频,1:回复评论)
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
