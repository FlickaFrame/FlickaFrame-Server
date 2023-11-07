package comment

import (
	"context"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/base"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
)

const (
	IsAuthor = "isAuthor" //视频的作者
)

// Comment 评论表
type Comment struct {
	base.Model
	VideoID       int64      `gorm:"column:video_id;not null;comment:视频id"`
	Content       string     `gorm:"column:content;type:varchar(1000);not null;comment:内容"`
	UserID        int64      `gorm:"column:user_id;not null;comment:评论人id"`
	LikeCount     int        `gorm:"column:like_count;default:0;comment:点赞数"`
	IsPublisher   bool       `gorm:"column:is_publisher;default:false;comment:是否为发布者"`
	ChildComments []*Comment `gorm:"-"` // 二级评论
	ParentID      int64      `gorm:"column:parent_id;default:null;comment:父评论id，也即第一级评论"`
	ParentComment *Comment   `gorm:"-"` // 父评论信息
	ReplyID       int64      `gorm:"column:reply_id;default:null;comment:被回复的评论id（没有则是回复父级评论，有则是回复这个人的评论）"`
}

func (*Comment) TableName() string {
	return "comment"
}

type Model struct {
	db *orm.DB
}

func NewCommentModel(db *orm.DB) *Model {
	return &Model{db}
}

// CreateParentComment 创建视频的一级评论
func (m *Model) CreateParentComment(ctx context.Context, doer, videoId int64, content string) (*Comment, error) {
	comment := Comment{
		Model:   base.NewModel(),
		Content: content,
		VideoID: videoId,
		UserID:  doer,
	}
	return &comment, m.db.WithContext(ctx).
		Create(&comment).Error
}

// CreateChildComment 创建视频的二级评论/回复评论
func (m *Model) CreateChildComment(ctx context.Context, doer, videoId int64, content string, parentCommentId, targetCommentId int64) (*Comment, error) {
	comm := Comment{
		Model:    base.NewModel(),
		Content:  content,
		VideoID:  videoId,
		UserID:   doer,
		ParentID: parentCommentId,
	}
	if targetCommentId != 0 {
		comm.ReplyID = targetCommentId
	}
	return &comm, m.db.WithContext(ctx).
		Create(&comm).Error
}

// FindOneComment 根据主键找评论
func (m *Model) FindOneComment(ctx context.Context, commentId int64) (*Comment, error) {
	var comment Comment
	err := m.db.WithContext(ctx).
		First(&comment, commentId).Error
	return &comment, err
}

// FindParentCommentByVideoId 根据VideoId查找一级评论
func (m *Model) FindParentCommentByVideoId(ctx context.Context, videoId int64, opts *CommentOption) ([]*Comment, error) {
	var comments []*Comment
	session := m.db.WithContext(ctx)
	session = orm.SetSessionPagination(session, opts)
	err := session.Where("video_id = ? and parent_id is null", videoId).Find(&comments).Error
	return comments, err
}

// FindCommentByUserId 根据用户Id查看评论
func (m *Model) FindCommentByUserId(ctx context.Context, userId int64, opts *CommentOption) (comments []*Comment, err error) {
	session := m.db.WithContext(ctx)
	session = orm.SetSessionPagination(session, opts)
	err = session.Where("user_id", userId).Find(&comments).Error
	return comments, err
}

// FindChildCommentByCommentId 根据评论Id找二级评论
func (m *Model) FindChildCommentByCommentId(ctx context.Context, parentId int64, opts *CommentOption) (comments []*Comment, err error) {
	session := m.db.WithContext(ctx)
	session = orm.SetSessionPagination(session, opts)
	err = session.Where("parent_id = ?", parentId).Find(&comments).Error
	return
}

func (m *Model) CountCommentByParentCommentId(ctx context.Context, parentCommentId int64) (int64, error) {
	var cnt int64
	return cnt, m.db.WithContext(ctx).
		Model(&Comment{}).
		Select("id").
		Where("parent_id = ?", parentCommentId).Count(&cnt).Error
}

func (m *Model) DeleteCommentByDoer(ctx context.Context, commentId, doerId int64) error {
	comment := &Comment{}
	comment.ID = commentId
	rowsAffected := m.db.WithContext(ctx).
		Where("user_id = ? and id = ?", doerId, commentId).
		Delete(comment).RowsAffected
	if rowsAffected == 0 {
		return fmt.Errorf("delete child comment fail")
	}
	return nil
}

func (m *Model) DeleteComment(ctx context.Context, commentId int64) error {
	comment := &Comment{}
	comment.ID = commentId
	rowsAffected := m.db.WithContext(ctx).
		Where("id = ?", commentId).
		Delete(comment).RowsAffected
	if rowsAffected == 0 {
		return fmt.Errorf("delete comment fail")
	}
	return nil
}

type CommentOption struct {
	orm.Paginator
	VideoID     int64  // 视频ID
	ParentId    int64  // 父评论ID
	UserId      int64  // 用户ID
	Limit       int    // 单次查询个数
	LimitOffset int    // 查询偏移量
	sortType    string // 排序方式(0:按时间倒序,1:按热度倒序)
}

// applySorts 排序规则
func (m *Model) applySorts(session *gorm.DB, sortType string) *gorm.DB {
	switch sortType {
	default:
		session = session.Order("`comment`.created_at desc")
	}
	return session
}

func (m *Model) list(ctx context.Context, opts *CommentOption) (comments []*Comment, err error) {
	session := m.db.WithContext(ctx)
	session = orm.SetSessionPagination(session, opts)
	if opts.VideoID != 0 {
		session = session.Where("video_id = ?", opts.VideoID)
	}
	if opts.ParentId != 0 {
		session = session.Where("parent_id = ?", opts.ParentId)
	}
	if opts.UserId != 0 {
		session = session.Where("user_id = ?", opts.UserId)
	}
	if opts.sortType != "" {
		session = session.Order("created_at desc")
	}
	err = session.Find(&comments).Error
	return
}
