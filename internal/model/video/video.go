package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"gorm.io/gorm"
	"time"
)

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

// FindByAuthorID 通过作者ID获取记录
func (m *VideoModel) FindByAuthorID(ctx context.Context, authorID int64, limit int) ([]*Video, error) {
	var result []*Video
	err := m.db.WithContext(ctx).
		Where("author_id = ?", authorID).
		Order("id desc").
		Limit(limit).
		Find(&result).Error
	return result, err
}

// FindByLatestTime 通过时间点来获取比该时间点早的十个视频
func (m *VideoModel) FindByLatestTime(ctx context.Context, latestTime time.Time, limit int) ([]*Video, error) {
	var result []*Video
	err := m.db.WithContext(ctx).
		Where("created_at < ?", latestTime).
		Order("created_at desc").
		Limit(limit).
		Find(&result).Error
	return result, err
}

//// GetVideoCountByAuthorID 通过作者 ID 获取作品数量
//func (m *VideoModel) GetVideoCountByAuthorID(authorID int64) (count int64, err error) {
//	m := query.Video
//	count, err = m.Where(m.AuthorID.Eq(authorID)).Count()
//	return count, err
//}
//
//// GetVideoIDByAuthorID 通过作者ID获取视频ID表
//func (m *VideoModel) GetVideoIDByAuthorID(authorID int64) (id []int64, err error) {
//	videos, err := GetVideoByAuthorID(authorID)
//	for _, video := range videos {
//		id = append(id, video.AuthorID)
//	}
//	return id, err
//}
//
//// DeleteVideoByID 通过记录ID删除对应记录
//func (m *VideoModel) DeleteVideoByID(id int64) (err error) {
//	m := query.Video
//	_, err = m.Where(m.ID.Eq(id)).Delete()
//	return err
//}
//
//// GetVideosByLatestTimeOrderByDESC 通过时间点来获取比该时间点早的十个视频
//func (m *VideoModel) GetVideosByLatestTimeOrderByDESC(latestTime time.Time) ([]*model.Video, error) {
//	m := query.Video
//	videos, err := m.Where(m.CreatedAt.Lt(latestTime)).Order(m.CreatedAt.Desc()).Limit(10).Find()
//	if err != nil {
//		log.Logger.Error(err)
//		return nil, err
//	}
//	return videos, nil
//}
//
//// GetVideosByAuthorIDAnTimeOrderByDESC 通过 AuthorID 和时间点来获取比该时间点早的十个视频
//func (m *VideoModel) GetVideosByAuthorIDAnTimeOrderByDESC(authorID int64, latestTime time.Time) ([]*model.Video, error) {
//	m := query.Video
//	videos, err := m.Where(m.AuthorID.Eq(authorID), m.CreatedAt.Lt(latestTime)).Limit(10).Find()
//	if err != nil {
//		log.Logger.Error(err)
//		return nil, err
//	}
//	return videos, nil
//}
//
//// GetVideosByAuthorID 通过作者 ID 获取视频列表
//func (m *VideoModel) GetVideosByAuthorID(authorID int64) ([]*model.Video, error) {
//	m := query.Video
//	videos, err := m.Where(m.AuthorID.Eq(authorID)).Find()
//	if err != nil {
//		log.Logger.Error(err)
//		return nil, err
//	}
//	return videos, nil
//}
