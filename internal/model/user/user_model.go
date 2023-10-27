package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
)

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

// GetUserFollowers returns range of user's followers.
func (m *UserModel) GetUserFollowers(ctx context.Context, userId uint, listOptions orm.ListOptions) ([]*User, error) {
	sess := m.db.WithContext(ctx).
		Select("`user`.*").
		Joins("LEFT", "follow", "`user`.id=follow.user_id").
		Where("follow.follow_id=?", userId)

	if listOptions.Page != 0 {
		sess = orm.SetSessionPagination(sess, &listOptions)
		users := make([]*User, 0, listOptions.PageSize)
		return users, sess.Find(&users).Error
	}

	users := make([]*User, 0, 8)
	err := sess.Find(&users).Error
	return users, err
}

// GetUserFollowing returns range of user's following.
func (m *UserModel) GetUserFollowing(ctx context.Context, userId uint, listOptions orm.ListOptions) ([]*User, error) {
	sess := m.db.WithContext(ctx).
		Select("`user`.*").
		Joins("LEFT", "follow", "`user`.id=follow.follow_id").
		Where("follow.user_id=?", userId)

	if listOptions.Page != 0 {
		sess = orm.SetSessionPagination(sess, &listOptions)
		users := make([]*User, 0, listOptions.PageSize)
		return users, sess.Find(&users).Error
	}

	users := make([]*User, 0, 8)
	return users, sess.Find(&users).Error
}
