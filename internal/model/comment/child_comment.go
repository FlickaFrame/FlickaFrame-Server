package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/snowflake"
	"time"
)

// ChildComment 二级评论表(评论一级评论)
type ChildComment struct {
	BasicComment
	ParentID      int64          `gorm:"column:parent_id;not null;comment:父评论id，也即第一级评论"`
	ParentComment *ParentComment `gorm:"-"` // 父评论信息
	ReplyID       int64          `gorm:"column:reply_id;default:null;comment:被回复的评论id（没有则是回复父级评论，有则是回复这个人的评论）"`
	LikeCount     int            `gorm:"column:like_count;default:0;comment:点赞数"`
	IsPublisher   bool           `gorm:"column:is_publisher;default:false;comment:是否为发布者"`
}

func (ChildComment) TableName() string {
	return "child_comment"
}

func (m *CommentModel) ListChildComment(ctx context.Context, parentId int64, opts Option) (comments []*ChildComment, err error) {
	session := m.applyOption(ctx, opts)
	err = session.Where("parent_id = ?", parentId).Find(&comments).Error
	return
}

// CreateChildComment 创建视频的二级评论
func (m *CommentModel) CreateChildComment(ctx context.Context, doer int64, videoId int64, content string, parentCommentId, targetCommentId int64) (*ChildComment, error) {
	comment := ChildComment{
		BasicComment: BasicComment{
			ID:        snowflake.CommentIDNode.Generate().Int64(),
			Content:   content,
			VideoID:   videoId,
			UserID:    doer,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ParentID: parentCommentId,
	}
	return &comment, m.db.WithContext(ctx).
		Create(&comment).Error
}
