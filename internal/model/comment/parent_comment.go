package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/snowflake"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"time"
)

// ParentComment 一级评论表(直接评论视频)
type ParentComment struct {
	BasicComment
	LikeCount     int             `gorm:"column:like_count;default:0;comment:点赞数"`
	IsPublisher   bool            `gorm:"column:is_publisher;default:false;comment:是否为发布者"`
	ChildComments []*ChildComment `gorm:"-"` // 二级评论
}

func (parent *ParentComment) TableName() string {
	return "parent_comment"
}

// LoadChildComments 加载二级评论
func (parent *ParentComment) LoadChildComments(ctx context.Context, db *orm.DB, opts Option) error {
	childComments, err := NewCommentModel(db).ListChildComment(ctx, parent.ID, opts)
	if err != nil {
		return err
	}
	parent.ChildComments = childComments
	return nil
}

func (m *CommentModel) ListParentComment(ctx context.Context, videoId int64, opts Option) (comments []*ParentComment, err error) {
	session := m.applyOption(ctx, opts)
	err = session.Where("video_id = ?", videoId).Find(&comments).Error
	return
}

// CreateParentComment 创建视频的一级评论
func (m *CommentModel) CreateParentComment(ctx context.Context, doer, videoId int64, content string) error {
	comment := ParentComment{
		BasicComment: BasicComment{
			ID:        snowflake.CommentIDNode.Generate().Int64(),
			Content:   content,
			VideoID:   videoId,
			UserID:    doer,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	return m.db.WithContext(ctx).
		Create(&comment).Error
}
