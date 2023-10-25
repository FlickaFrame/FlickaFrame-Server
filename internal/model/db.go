package model

import (
	follow_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/follow"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user_model.User{},
		&video_model.Video{},
		&follow_model.Follow{},
	)
}
