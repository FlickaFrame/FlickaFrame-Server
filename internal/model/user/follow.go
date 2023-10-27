package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

// Follow 关注
type Follow struct {
	gorm.Model
	UserID         uint  `gorm:"uniqueIndex:idx_follow"`
	User           *User `gorm:"-"`
	FollowedUserID uint  `gorm:"uniqueIndex:idx_follow"`
	FollowedUser   *User `gorm:"-"`
	Status         bool  `gorm:"index;not null"`
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

func (m *FollowModel) Insert(ctx context.Context, data *Follow) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *FollowModel) FindOne(ctx context.Context, id int64) (*Follow, error) {
	var result Follow
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *FollowModel) IsFollow(ctx context.Context, userId, followedUserId uint) (bool, error) {
	if userId == 0 || followedUserId == 0 {
		return false, nil
	}
	var result Follow
	err := m.db.WithContext(ctx).Where("user_id = ? AND followed_user_id = ?", userId, followedUserId).First(&result).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return result.Status, err
}

func (m *FollowModel) FollowOrUnFollow(ctx context.Context, userId, followedUserId uint, status bool) error {
	result := Follow{
		UserID:         userId,
		FollowedUserID: followedUserId,
		Status:         status,
	}
	err := m.db.WithContext(ctx).Where("user_id =? and followed_user_id=?", userId, followedUserId).FirstOrCreate(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	result.Status = status
	return m.db.WithContext(ctx).Save(result).Error
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
