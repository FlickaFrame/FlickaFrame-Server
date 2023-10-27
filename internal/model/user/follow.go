package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// Follow 关注
type Follow struct {
	gorm.Model
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
		if err := tx.Model(&User{}).Where("id = ?", followID).Update("num_followers", gorm.Expr("num_followers + ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", userId).Update("num_following", gorm.Expr("num_following + ?", 1)).Error; err != nil {
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
		if err := tx.Model(&User{}).Where("id = ?", followID).Update("num_followers", gorm.Expr("num_followers - ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", userID).Update("num_following", gorm.Expr("num_following - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
}

func (m *FollowModel) Update(ctx context.Context, data *Follow) error {
	return m.db.WithContext(ctx).Save(data).Error
}

func (m *FollowModel) UpdateFields(ctx context.Context, id int64, values map[string]interface{}) error {
	return m.db.WithContext(ctx).Model(&Follow{}).Where("id = ?", id).Updates(values).Error
}

func (m *FollowModel) FindByUserIDAndFollowedUserID(ctx context.Context, userId, followedUserId int64) (*Follow, error) {
	var result Follow
	err := m.db.WithContext(ctx).
		Where("user_id = ? AND followed_user_id = ?", userId, followedUserId).
		First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &result, err
}

func (m *FollowModel) FindByUserId(ctx context.Context, userId int64, limit int) ([]*Follow, error) {
	var result []*Follow
	err := m.db.WithContext(ctx).
		Where("user_id = ? AND follow_status = ?", userId, 1).
		Order("id desc").
		Limit(limit).
		Find(&result).Error

	return result, err
}

func (m *FollowModel) FindByFollowedUserIds(ctx context.Context, followedUserIds []int64) ([]*Follow, error) {
	var result []*Follow
	err := m.db.WithContext(ctx).
		Where("followed_user_id in (?)", followedUserIds).
		Find(&result).Error

	return result, err
}
