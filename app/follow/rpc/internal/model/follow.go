package model

import (
	"context"
	"errors"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
	"sort"
)

type Follow struct {
	orm.Model
	UserID         int64 `gorm:"uniqueIndex:idx_follow"`
	FollowedUserID int64 `gorm:"uniqueIndex:idx_follow"`
	FollowStatus   int   `gorm:"status"`
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
	return m.db.WithContext(ctx).
		Model(&Follow{}).
		Where("id = ?", id).
		Updates(values).
		Error
}

// FindByUserIDAndFollowedUserID 获取关注关系
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

func (m *FollowModel) FindByFollowedUserId(ctx context.Context, userId int64, limit int) ([]*Follow, error) {
	var result []*Follow
	err := m.db.WithContext(ctx).
		Where("followed_user_id = ? AND follow_status = ?", userId, 1).
		Order("id desc").
		Limit(limit).
		Find(&result).Error
	return result, err
}

// FindByFollowedUserIds 获取关注的人列表
func (m *FollowModel) FindByFollowedUserIds(ctx context.Context, followedUserIds []int64, userId int64) ([]*Follow, error) {
	var result []*Follow
	err := m.db.WithContext(ctx).
		Where("followed_user_id in (?)", followedUserIds).
		Where("user_id =?", userId).
		Find(&result).Error
	val2idx := make(map[int64]int, len(followedUserIds))
	for i, v := range followedUserIds {
		val2idx[v] = i
	}
	sort.Slice(result, func(i, j int) bool {
		return val2idx[result[i].FollowedUserID] < val2idx[result[j].FollowedUserID]
	})
	return result, err
}

func (m *FollowModel) FindByUserIds(ctx context.Context, userIds []int64, followedUserId int64) ([]*Follow, error) {
	var result []*Follow
	err := m.db.WithContext(ctx).
		Where("user_id in (?) ", userIds).
		Where("followed_user_id=?", followedUserId).
		Find(&result).Error
	val2idx := make(map[int64]int, len(userIds))
	for i, v := range userIds {
		val2idx[v] = i
	}
	sort.Slice(result, func(i, j int) bool {
		return val2idx[result[i].UserID] < val2idx[result[j].UserID]
	})
	return result, err
}
