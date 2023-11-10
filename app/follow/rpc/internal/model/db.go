package model

import "github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"

// Migrate  数据库迁移
func Migrate(db *orm.DB) error {
	return db.AutoMigrate(
		&Follow{},
		&FollowCount{},
	)
}
