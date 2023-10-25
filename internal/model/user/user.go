package user

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string
	AvtarUrl       string
	Password       string
	Phone          string `gorm:"uniqueIndex;not null;size:11"`
	FollowingCount int
	FollowerCount  int
}

func (u *User) TableName() string {
	return "user"
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *UserModel) MustFindOne(ctx context.Context, id int64) *User {
	user, _ := m.FindOne(ctx, id)
	return user
}
