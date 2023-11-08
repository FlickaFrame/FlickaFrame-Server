package favorite

import (
	"context"
	"errors"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/base"
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/comment"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
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

func (m *Model) DeleteVideoFavorite(ctx context.Context, userId, videoId int64) error {
	return m.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			rowsAffected := m.db.WithContext(ctx).
				Where("target_id = ? AND user_id = ?", videoId, userId).
				Delete(&Favorite{}).RowsAffected
			if rowsAffected == 0 {
				return fmt.Errorf("无法取消不存在的点赞")
			}
			if err := tx.Model(&comment_model.Comment{}).
				Where("id = ?", videoId).
				Update("favorite_count",
					gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
				return err
			}
			return nil
		})
}

func (m *Model) DeleteCommentFavorite(ctx context.Context, userId, commentId int64) error {
	return m.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			rowsAffected := m.db.WithContext(ctx).
				Where("target_id = ? AND user_id = ?", commentId, userId).
				Delete(&Favorite{}).RowsAffected
			if rowsAffected == 0 {
				return fmt.Errorf("无法取消不存在的点赞")
			}
			if err := tx.Model(&comment_model.Comment{}).
				Where("id = ?", commentId).
				Update("like_count",
					gorm.Expr("like_count - ?", 1)).Error; err != nil {
				return err
			}
			return nil
		})
}

func (m *Model) CreateCommentFavorite(ctx context.Context, userid, commentId int64) error {
	return m.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&Favorite{
				Model:    base.NewModel(),
				TargetID: commentId,
				UserID:   userid,
				Type:     CommentFavoriteType,
			}).Error; err != nil {
				return err
			}
			if err := tx.Model(&comment_model.Comment{}).
				Where("id = ?", commentId).
				Update("like_count",
					gorm.Expr("like_count + ?", 1)).Error; err != nil {
				return err
			}
			return nil
		})
}

func (m *Model) CreateVideoFavorite(ctx context.Context, userid, videoId int64) error {
	return m.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			if err := tx.Create(&Favorite{
				Model:    base.NewModel(),
				TargetID: videoId,
				UserID:   userid,
				Type:     VideoFavoriteType,
			}).Error; err != nil {
				return err
			}
			if err := tx.Model(&video_model.Video{}).
				Where("id = ?", videoId).
				Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
				return err
			}
			return nil
		})
}
