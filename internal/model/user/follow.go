package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// Follow 关注
type Follow struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID         uint  `gorm:"uniqueIndex:idx_follow"`
	User           *User `gorm:"-"`
	FollowedUserID uint  `gorm:"uniqueIndex:idx_follow"`
	FollowedUser   *User `gorm:"-"`
}

func (m *Follow) TableName() string {
	return "follow"
}

type FollowModel struct {
	db *gorm.DB
}

func NewFollowModel(db *gorm.DB) *FollowModel {
	return &FollowModel{
		db: db,
	}
}

// IsFollowing returns true if user is following followID.
func (m *UserModel) IsFollowing(ctx context.Context, userId, followID uint) bool {
	if userId == 0 || followID == 0 {
		return false
	}
	var result Follow
	err := m.db.WithContext(ctx).
		Where("user_id = ? AND followed_user_id = ?", userId, followID).
		First(&result).Error
	if err != nil {
		logx.Debugf("IsFollowing error: %v", err)
		return false
	}
	return true
}

// FollowUser marks someone be  follower.
func (m *UserModel) FollowUser(ctx context.Context, userId, followID uint) error {
	if userId == followID || m.IsFollowing(ctx, userId, followID) {
		return nil
	}
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&Follow{UserID: userId, FollowedUserID: followID}).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", followID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", userId).Update("following_count", gorm.Expr("following_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// UnfollowUser unmarks someone as another's follower.
func (m *UserModel) UnfollowUser(ctx context.Context, userID, followID uint) error {
	if userID == followID || !m.IsFollowing(ctx, userID, followID) {
		return nil
	}
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? AND followed_user_id = ?", userID, followID).Delete(&Follow{}).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", followID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", userID).Update("following_count", gorm.Expr("following_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetUserFollowers returns range of user's followers.
func (m *UserModel) GetUserFollowers(ctx context.Context, userId uint, listOptions orm.ListOptions) ([]*User, error) {
	sess := m.db.WithContext(ctx).
		Select("`user`.*").
		Joins("join follow on`user`.id=`follow`.user_id").
		Where("`follow`.followed_user_id=?", userId)

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
		Joins("join follow on `user`.id=`follow`.followed_user_id").
		Where("`follow`.user_id=?", userId)
	var users []*User
	if listOptions.Page != 0 {
		sess = orm.SetSessionPagination(sess, &listOptions)
		users = make([]*User, 0, listOptions.PageSize)
	} else {
		users = make([]*User, 0, 8)
	}

	return users, sess.Find(&users).Error
}

func (m *UserModel) CountFollowers(ctx context.Context, userId uint) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&Follow{}).Where("followed_user_id = ?", userId).Count(&count).Error
	return count, err
}

func (m *UserModel) CountFollowing(ctx context.Context, userId uint) (int64, error) {
	var count int64
	err := m.db.WithContext(ctx).Model(&Follow{}).Where("user_id = ?", userId).Count(&count).Error
	return count, err
}
