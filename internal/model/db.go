package model

import (
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/comment"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user_model.User{},
		&video_model.Video{},
		&video_model.Tag{},
		&video_model.VideoTag{},
		&video_model.Category{},
		&user_model.Follow{},
		&comment_model.Comment{},
	)
}
