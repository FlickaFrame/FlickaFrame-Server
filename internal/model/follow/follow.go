package follow

import (
	"context"
	"errors"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	Status         int
	UserID         int64            `gorm:"uniqueIndex:idx_follow"`
	User           *user_model.User `gorm:"-"`
	FollowedUserID int64            `gorm:"uniqueIndex:idx_follow"`
	FollowedUser   *user_model.User `gorm:"-"`
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
