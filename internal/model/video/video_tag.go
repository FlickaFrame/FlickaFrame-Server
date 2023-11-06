package video

import (
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/base"
)

type Tag struct {
	base.Model

	Name string `gorm:"unique"`
}

func (t *Tag) TableName() string {
	return "tags"
}

type VideoTag struct {
	base.Model

	VideoID int64  `gorm:"index"`
	Video   *Video `gorm:"-"`

	TagID int64 `gorm:"index"`
	Tag   *Tag  `gorm:"-"`
}

func (vt *VideoTag) TableName() string {
	return "video_tags"
}
