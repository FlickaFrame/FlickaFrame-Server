package comment

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/snowflake"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
	"time"
)

const DefaultLimit = 10

type Comment struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Content     string          `gorm:"not null"`                    // 评论内容
	VideoID     int64           `gorm:"index:idx_videoid;not null"`  // 视频ID
	ParentID    int64           `gorm:"index:idx_parentid;"`         // 父评论ID
	TargetID    int64           `gorm:"index:idx_targetid;"`         // 回复对应评论ID
	OwnerUID    int64           `gorm:"index:idx_owneruid;not null"` // 评论者ID
	User        user_model.User `gorm:"-"`                           // 评论者信息
	ToUser      user_model.User `gorm:"-"`                           // 被回复者信息
	ReplayCount int             `gorm:"default:0"`                   // 回复数
	LikeCount   int             `gorm:"default:0"`                   // 点赞数
}

func (c *Comment) TableName() string {
	return "comment"
}

type CommentModel struct {
	db *orm.DB
}

func NewCommentModel(db *orm.DB) *CommentModel {
	return &CommentModel{db}
}

type Option struct {
	VideoID     int64 // 视频ID
	ParentId    int   // 父评论ID
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

func (m *CommentModel) List(ctx context.Context, opts Option) ([]*Comment, error) {
	var result []*Comment
	err := m.applyOption(ctx, opts).Find(&result).Error
	return result, err
}

// CreateVideoComment 创建视频的一级评论
func (m *CommentModel) CreateVideoComment(ctx context.Context, doer, videoId int64, content string) error {
	comment := Comment{
		ID:       snowflake.CommentIDNode.Generate().Int64(),
		Content:  content,
		VideoID:  videoId,
		OwnerUID: doer,
	}
	return m.db.WithContext(ctx).
		Create(&comment).Error
}

// CreateReplyComment 创建视频的二级评论（评论的回复）
func (m *CommentModel) CreateReplyComment(ctx context.Context, doer int64, videoId int64, content string, parentId int64, targetId int64) error {
	comment := Comment{
		ID:       snowflake.CommentIDNode.Generate().Int64(),
		Content:  content,
		VideoID:  videoId,
		OwnerUID: doer,
		ParentID: parentId,
		TargetID: targetId,
	}
	return m.db.WithContext(ctx).
		Create(&comment).Error
}

func (m *CommentModel) Insert(ctx context.Context, comment Comment) error {
	return m.db.WithContext(ctx).
		Create(&comment).Error
}

func (m *CommentModel) Delete(ctx context.Context, id int64) error {
	// TODO删除评论时,需要删除对应的回复
	return m.db.WithContext(ctx).
		Delete(&Comment{}, id).Error
}

func (m *CommentModel) Update(ctx context.Context, id int64, content string) error {
	return m.db.WithContext(ctx).
		Model(&Comment{}).
		Where("id = ?", id).
		Update("content", content).Error
}

func (m *CommentModel) FindOne(ctx context.Context, id int64) (*Comment, error) {
	var comment Comment
	err := m.db.WithContext(ctx).First(&comment, id).Error
	return &comment, err
}
