package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"gorm.io/gorm"
	"time"
)

const DefaultLimit = 10

type Video struct {
	gorm.Model
	Title    string
	ETag     string
	PlayUrl  string
	CoverUrl string

	FavoriteCount int
	CommentCount  int

	AuthorID uint       `gorm:"index"`
	Author   *user.User `gorm:"-"`

	CategoryID uint `gorm:"index"`
}

func (v *Video) TableName() string {
	return "video"
}

func (v *Video) loadAuthor(ctx context.Context, db *gorm.DB) error {
	var author user.User
	err := db.WithContext(ctx).
		Where("id = ?", v.AuthorID).
		First(&author).Error
	v.Author = &author
	return err
}

type VideoModelInterface interface {
	Insert(ctx context.Context, video *Video) error
	FindOne(ctx context.Context, id int64) (*Video, error)
	List(ctx context.Context, opts ListOption) ([]*Video, error)
	Count(ctx context.Context, opts ListOption) (count int64, err error)
}

var _ VideoModelInterface = (*VideoModel)(nil)

type VideoModel struct {
	db *gorm.DB
}

func NewVideoModel(db *gorm.DB) *VideoModel {
	return &VideoModel{
		db: db,
	}
}

// Insert 创建视频记录
func (m *VideoModel) Insert(ctx context.Context, video *Video) error {
	return m.db.WithContext(ctx).Create(video).Error
}

// FindOne 通过视频ID获取对应记录
func (m *VideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	var result Video
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

// ListOption 查找选项
type ListOption struct {
	AuthorID   uint      // 作者ID
	LatestTime time.Time // 最新时间(分页)
	Limit      int       // 限制数量(分页)
	QueryAll   bool      // 是否查询所有(分页)
	CategoryID uint      // 分类ID
}

func (m *VideoModel) applyOption(ctx context.Context, opts ListOption) *gorm.DB {
	session := m.db.WithContext(ctx)
	// 根据作者ID
	if opts.AuthorID != 0 {
		session = session.Where("author_id = ?", opts.AuthorID)
	}
	// 根据分类ID
	if opts.CategoryID != 0 {
		session = session.Where("category_id = ?", opts.CategoryID)
	}
	// 分页
	if opts.Limit == 0 {
		opts.Limit = DefaultLimit
	}
	if !opts.QueryAll {
		session = session.Where("created_at <= ?", opts.LatestTime)
		session = session.Limit(opts.Limit)
	}
	return session.Order("created_at desc")
}

// List 通过时间点来获取比该时间点早的十个视频
func (m *VideoModel) List(ctx context.Context, opts ListOption) ([]*Video, error) {
	var result []*Video
	err := m.applyOption(ctx, opts).Find(&result).Error
	for _, v := range result {
		err = v.loadAuthor(ctx, m.db)
		if err != nil {
			return nil, err
		}
	}
	return result, err
}

// Count 通过作者 ID 获取视频数量
func (m *VideoModel) Count(ctx context.Context, opts ListOption) (count int64, err error) {
	err = m.applyOption(ctx, opts).Count(&count).Error
	return
}

func (m *VideoModel) ListVideoByAuthorId(ctx context.Context, authorId uint) ([]*Video, error) {
	var result []*Video
	err := m.db.WithContext(ctx).Where("author_id = ?", authorId).Find(&result).Error
	return result, err
}

func (m *VideoModel) ListVideoByCategoryId(ctx context.Context, categoryId uint) ([]*Video, error) {
	var result []*Video
	err := m.db.WithContext(ctx).Where("category_id = ?", categoryId).Find(&result).Error
	return result, err
}
