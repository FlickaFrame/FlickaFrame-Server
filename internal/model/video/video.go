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

	AuthorID int64      `gorm:"index"`
	Author   *user.User `gorm:"-"`

	CategoryID int64 `gorm:"index"`
}

func (v *Video) TableName() string {
	return "video"
}

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

// Option 查找选项
type Option struct {
	AuthorID   uint      // 作者ID
	LatestTime time.Time // 最新时间(分页)
	Limit      int       // 限制数量(分页)
	QueryAll   bool      // 是否查询所有(分页)
}

func (m *VideoModel) applyOption(ctx context.Context, opts Option) *gorm.DB {
	session := m.db.WithContext(ctx)
	if opts.Limit == 0 {
		opts.Limit = DefaultLimit
	}
	if opts.AuthorID != 0 {
		session = session.Where("author_id = ?", opts.AuthorID)
	}

	if !opts.QueryAll {
		session = session.Where("created_at <= ?", opts.LatestTime)
		session = session.Limit(opts.Limit)
	}
	return session.Order("created_at desc")
}

// List 通过时间点来获取比该时间点早的十个视频
func (m *VideoModel) List(ctx context.Context, opts Option) ([]*Video, error) {
	var result []*Video
	err := m.applyOption(ctx, opts).Find(&result).Error
	return result, err
}

// Count 通过作者 ID 获取视频数量
func (m *VideoModel) Count(ctx context.Context, opts Option) (count int64, err error) {
	err = m.applyOption(ctx, opts).Count(&count).Error
	return
}
