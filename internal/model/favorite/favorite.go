package favorite

import (
	"context"
	"errors"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/base"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	cacheFlickaFrameFavoritePrefix = "cache:flickaFrame:favorite:user:"
)

type FavoriteType = int

const (
	VideoFavoriteType = iota + 1
	CommentFavoriteType
)

// Favorite 点赞
type Favorite struct {
	base.Model
	UserID   int64 `gorm:"uniqueIndex:idx_target_user;not null"`
	TargetID int64 `gorm:"uniqueIndex:idx_target_user;not null"`
	Type     int
}

type Model struct {
	db *orm.DB
}

func NewFavoriteModel(db *orm.DB) *Model {
	return &Model{db}
}

func (m *Model) IsExist(ctx context.Context, targetId, userId int64) (bool, error) {
	err := m.db.WithContext(ctx).
		Where("target_id = ? AND user_id = ?", targetId, userId).
		First(&Favorite{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

func (m *Model) Delete(ctx context.Context, userId, targetId int64) error {
	rowsAffected := m.db.WithContext(ctx).
		Where("target_id = ? AND user_id = ?", targetId, userId).
		Delete(&Favorite{}).RowsAffected
	if rowsAffected == 0 {
		return fmt.Errorf("无法取消不存在的点赞")
	}
	return nil
}

func (m *Model) Create(ctx context.Context, userId, targetId int64, typ int) error {
	result := Favorite{
		Model:    base.NewModel(),
		TargetID: targetId,
		UserID:   userId,
		Type:     typ,
	}
	err := m.db.WithContext(ctx).Create(&result).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return code.DuplicateFavoriteErr
	}
	return err
}
