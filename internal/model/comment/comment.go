package comment

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"gorm.io/gorm"
)

const DefaultLimit = 10

type Comment struct {
	gorm.Model
	Content     string          `gorm:"not null"`                    // 评论内容
	VideoID     uint            `gorm:"index:idx_videoid;not null"`  // 视频ID
	ParentID    uint            `gorm:"index:idx_parentid;"`         // 父评论ID
	TargetID    uint            `gorm:"index:idx_targetid;"`         // 回复对应评论ID
	OwnerUID    uint            `gorm:"index:idx_owneruid;not null"` // 评论者ID
	User        user_model.User `gorm:"-"`                           // 评论者信息
	ToUser      user_model.User `gorm:"-"`                           // 被回复者信息
	ReplayCount int             `gorm:"default:0"`                   // 回复数
	LikeCount   int             `gorm:"default:0"`                   // 点赞数
}

func (c *Comment) TableName() string {
	return "comment"
}

type CommentModel struct {
	db *gorm.DB
}

func NewCommentModel(db *gorm.DB) *CommentModel {
	return &CommentModel{db}
}

type Option struct {
	VideoID     uint // 视频ID
	ParentId    int  // 父评论ID
	Limit       int  // 单次查询个数
	LimitOffset int  // 查询偏移量
	OrderBy     int  // 排序方式(0:按时间倒序,1:按热度倒序)
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

func (m *CommentModel) CreateVideoComment(ctx context.Context, doer uint, videoId uint, content string) error {
	comment := Comment{
		Content:  content,
		VideoID:  videoId,
		OwnerUID: doer,
	}
	return m.db.WithContext(ctx).
		Create(&comment).Error
}

func (m *CommentModel) Insert(ctx context.Context, comment Comment) error {
	return m.db.WithContext(ctx).
		Create(&comment).Error
}

func (m *CommentModel) Delete(ctx context.Context, id uint) error {
	// TODO删除评论时,需要删除对应的回复
	return m.db.WithContext(ctx).
		Delete(&Comment{}, id).Error
}

func (m *CommentModel) Update(ctx context.Context, id uint, content string) error {
	return m.db.WithContext(ctx).
		Model(&Comment{}).
		Where("id = ?", id).
		Update("content", content).Error
}

func (m *CommentModel) FindOne(ctx context.Context, id uint) (*Comment, error) {
	var comment Comment
	err := m.db.WithContext(ctx).First(&comment, id).Error
	return &comment, err
}
