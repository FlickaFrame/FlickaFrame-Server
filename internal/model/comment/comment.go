package comment

import (
	"context"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/base"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
)

const DefaultLimit = 10
const (
	IsAuthor = "isAuthor"
)

type BasicComment struct {
	base.Model
	VideoID int64            `gorm:"column:video_id;not null;comment:视频id"`
	Content string           `gorm:"column:content;type:varchar(1000);not null;comment:内容"`
	UserID  int64            `gorm:"column:user_id;not null;comment:评论人id"`
	User    *user_model.User `gorm:"-"` // 评论人信息
}

type CommentModel struct {
	db *orm.DB
}

func NewCommentModel(db *orm.DB) *CommentModel {
	return &CommentModel{db}
}

type Option struct {
	VideoID     int64 // 视频ID
	ParentId    int64 // 父评论ID
	Limit       int   // 单次查询个数
	LimitOffset int   // 查询偏移量
	OrderBy     int   // 排序方式(0:按时间倒序,1:按热度倒序)
}

func (m *CommentModel) applyOption(ctx context.Context, opts Option) *gorm.DB {
	session := m.db.WithContext(ctx)
	if opts.VideoID != 0 {
		session = session.Where("video_id = ?", opts.VideoID)
	}
	if opts.ParentId != 0 {
		session = session.Where("parent_id = ?", opts.ParentId)
	}
	if opts.Limit == 0 {
		opts.Limit = DefaultLimit
	}
	if opts.LimitOffset != 0 {
		session = session.Offset(opts.LimitOffset)
	}
	if opts.OrderBy == 0 {
		session = session.Order("created_at desc")
	}
	return session
}

// ListParentComments 获取视频的一级评论
func (m *CommentModel) ListParentComments(ctx context.Context, videoId int64) []*ParentComment {
	return nil
}

// ListChildComments 获取视频的二级评论
func (m *CommentModel) ListChildComments(ctx context.Context, commentId int64) []*ChildComment {
	return nil
}

func (m *CommentModel) FindParentComment(ctx context.Context, id int64) (*ParentComment, error) {
	var comment ParentComment
	err := m.db.WithContext(ctx).First(&comment, id).Error
	return &comment, err
}

func (m *CommentModel) FindChildComment(ctx context.Context, id int64) (*ChildComment, error) {
	var comment ChildComment
	err := m.db.WithContext(ctx).First(&comment, id).Error
	return &comment, err
}

func (m *CommentModel) DeleteParentComment(ctx context.Context, id, doerId int64) error {
	comment := &ParentComment{}
	comment.ID = id
	rowsAffected := m.db.WithContext(ctx).
		Where("user_id = ? and id = ?", doerId, id).
		Delete(comment).RowsAffected
	if rowsAffected == 0 {
		return fmt.Errorf("delete parent comment fail")
	}
	return nil
}

func (m *CommentModel) DeleteChildComment(ctx context.Context, id, doerId int64) error {
	comment := &ChildComment{}
	comment.ID = id
	rowsAffected := m.db.WithContext(ctx).
		Where("user_id = ? and id = ?", doerId, id).
		Delete(comment).RowsAffected
	if rowsAffected == 0 {
		return fmt.Errorf("delete child comment fail")
	}
	return nil
}

//func (m *CommentModel) ListCommentTag(ctx context.Context, id int64) {
//
//}
