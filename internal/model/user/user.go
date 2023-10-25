package user

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName       string
	AvtarUrl       string
	Sex            int64
	Password       string
	Phone          string
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

func (m *UserModel) Insert(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Create(data).Error
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

func (m *UserModel) FindOneByPhone(ctx context.Context, phone string) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).Where("phone = ?", phone).First(&result).Error
	return &result, err
}
