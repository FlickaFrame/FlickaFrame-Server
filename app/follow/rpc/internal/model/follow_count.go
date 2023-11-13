package model

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
	"sort"
)

type FollowCount struct {
	orm.Model
	UserID      int64 `gorm:"uniqueIndex:idx_follow_count"`
	FollowCount int   `gorm:"follow_count"`
	FansCount   int   `gorm:"fans_count"`
}

func (m *FollowCount) TableName() string {
	return "follow_count"
}

type FollowCountModel struct {
	db *gorm.DB
}

func NewFollowCountModel(db *gorm.DB) *FollowCountModel {
	return &FollowCountModel{
		db: db,
	}
}

func (m *FollowCountModel) Insert(ctx context.Context, data *FollowCount) error {
	return m.db.Create(data).Error
}

func (m *FollowCountModel) FindOne(ctx context.Context, id int64) (*FollowCount, error) {
	var result FollowCount
	err := m.db.Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *FollowCountModel) Update(ctx context.Context, data *FollowCount) error {
	return m.db.Save(data).Error
}

func (m *FollowCountModel) IncrFollowCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("INSERT INTO follow_count (user_id, follow_count) VALUES (?, 1) ON DUPLICATE KEY UPDATE follow_count = follow_count + 1", userId).
		Error
}

func (m *FollowCountModel) DecrFollowCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("UPDATE follow_count SET follow_count = follow_count - 1 WHERE user_id = ? AND follow_count > 0", userId).
		Error
}

func (m *FollowCountModel) IncrFansCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("INSERT INTO follow_count (user_id, fans_count) VALUES (?, 1) ON DUPLICATE KEY UPDATE fans_count = fans_count + 1", userId).
		Error
}

func (m *FollowCountModel) DecrFansCount(ctx context.Context, userId int64) error {
	return m.db.WithContext(ctx).
		Exec("UPDATE follow_count SET fans_count = fans_count - 1 WHERE user_id = ? AND fans_count > 0", userId).
		Error
}

func (m *FollowCountModel) FindByUserIds(ctx context.Context, userIds []int64) ([]*FollowCount, error) {
	var result []*FollowCount
	err := m.db.WithContext(ctx).Where("user_id IN ?", userIds).Find(&result).Error
	val2idx := make(map[int64]int, len(userIds))
	for i, v := range userIds {
		val2idx[v] = i
	}
	sort.Slice(result, func(i, j int) bool {
		return val2idx[result[i].UserID] < val2idx[result[j].UserID]
	})
	return result, err
}
