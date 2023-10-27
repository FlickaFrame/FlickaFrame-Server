package user

import (
	"context"
	"gorm.io/gorm"
)

const (
	SexUnknown = iota
	SexMale
	SexFemale
)

type User struct {
	gorm.Model
	NickName  string `gorm:"type:varchar(32)"` // 昵称
	AvatarUrl string // 头像地址
	Age       int    // 年龄
	Gender    int    `gorm:"type:enum('0','1','2')"` // 性别
	Password  string // 密码
	Phone     string `gorm:"type:varchar(100);index:idx_phone,unique"` // 手机号
	Slogan    string // 个人简介
	TikTokID  string `gorm:"type:varchar(100);index:idx_tiktok,unique"` // 抖音ID

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

func (m *UserModel) FindOne(ctx context.Context, id uint) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *UserModel) MustFindOne(ctx context.Context, id uint) *User {
	user, _ := m.FindOne(ctx, id)
	return user
}

func (m *UserModel) FindOneByPhone(ctx context.Context, phone string) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).Where("phone = ?", phone).First(&result).Error
	return &result, err
}
