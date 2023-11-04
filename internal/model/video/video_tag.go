package video

import (
	"time"
)

type Tag struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name string `gorm:"unique"`
}

func (t *Tag) TableName() string {
	return "tags"
}

type VideoTag struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	VideoID int64  `gorm:"index"`
	Video   *Video `gorm:"-"`

	TagID int64 `gorm:"index"`
	Tag   *Tag  `gorm:"-"`
}

func (vt *VideoTag) TableName() string {
	return "video_tags"
}
