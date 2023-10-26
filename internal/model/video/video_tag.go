package video

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type VideoTag struct {
	gorm.Model

	VideoID int64 `gorm:"index"`
	Video   *Video

	TagID int64 `gorm:"index"`
	Tag   *Tag
}
