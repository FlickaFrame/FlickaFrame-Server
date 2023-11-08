package model

import (
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/comment"
	favorite_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/favorite"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/notice"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/user"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
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
		&favorite_model.Favorite{},
		&notice.Notice{},
	)
}
