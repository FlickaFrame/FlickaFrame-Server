package video

import (
	"context"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/base"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/snowflake"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sort"
	"time"
)

const DefaultLimit = 10

var cacheFlickaFrameVideoIdPrefix = "cache:flickaFrame:video:id:"

type Video struct {
	base.Model
	Title         string     `gorm:"column:title;comment:视频标题"`
	PlayUrl       string     `gorm:"column:play_url;comment:播放地址"`
	ThumbUrl      string     `gorm:"column:thumb_url"`
	FavoriteCount int64      `gorm:"default:0"` // 点赞数量
	CommentCount  int64      `gorm:"default:0"` // 评论数量
	AuthorID      int64      `gorm:"index"`     // 作者ID
	Author        *user.User `gorm:"-"`         // 作者
	CategoryID    int64      `gorm:"index"`     // 分类ID
	Category      *Category  `gorm:"-"`         // 分类 `gorm:"-"`
	Tags          []*Tag     `gorm:"-"`
	Description   string     // 视频描述
	PublishTime   time.Time  // 发布时间
	PublishStatus int        `gorm:"default:0"` // 发布状态 0:未发布 1:已发布
	Visibility    int        `gorm:"default:0"` // 可见性 0:公开 1:私有
	VideoDuration float32    // 视频时长
	VideoHeight   float32    // 视频高度
	VideoWidth    float32    // 视频宽度
}

func (v *Video) TableName() string {
	return "video"
}

func (v *Video) LoadCategory(ctx context.Context, db *orm.DB) error {
	var category Category
	err := db.WithContext(ctx).
		Where("id = ?", v.CategoryID).
		First(&category).Error
	v.Category = &category
	return err
}

func (v *Video) LoadTags(ctx context.Context, db *orm.DB) error {
	var tags []*Tag
	err := db.WithContext(ctx).
		Select("tags.*").
		Joins("LEFT JOIN video_tags ON video_tags.tag_id = tags.id").
		Where("video_id = ?", v.ID).
		Find(&tags).Error
	v.Tags = tags
	return err
}

func (v *Video) LoadAttributes(ctx context.Context, db *orm.DB) error {
	err := v.LoadCategory(ctx, db)
	if err != nil {
		return err
	}
	err = v.LoadTags(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

type VideoModel struct {
	db *orm.DB
}

func NewVideoModel(db *orm.DB) *VideoModel {
	return &VideoModel{
		db: db,
	}
}

// Insert 创建视频记录
func (m *VideoModel) Insert(ctx context.Context, video *Video) error {
	video.ID = snowflake.VideoIDNode.Generate().Int64()
	return m.db.WithContext(ctx).Create(video).Error
}

// FindOne 通过视频ID获取对应记录
func (m *VideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	var result Video
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *VideoModel) FindByIDsAndCategory(ctx context.Context, ids []int64, category int64) ([]*Video, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var result []*Video
	sess := m.db.WithContext(ctx).Where("id IN ?", ids)
	if category != 0 { // 0 表示不限制分类
		sess = sess.Where("category_id = ?", category)
	}
	err := sess.Find(&result).Error
	val2idx := make(map[int64]int, len(result))
	for i, v := range ids {
		val2idx[v] = i
	}
	sort.Slice(result, func(i, j int) bool {
		return val2idx[result[i].ID] < val2idx[result[j].ID]
	})
	return result, err
}

// FindByIDs 通过视频ID列表获取对应记录
func (m *VideoModel) FindByIDs(ctx context.Context, ids []int64) ([]*Video, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	var result []*Video
	err := m.db.WithContext(ctx).Where("id IN ?", ids).Find(&result).Error
	val2idx := make(map[int64]int, len(result))
	for i, v := range ids {
		val2idx[v] = i
	}
	sort.Slice(result, func(i, j int) bool {
		return val2idx[result[i].ID] < val2idx[result[j].ID]
	})
	return result, err
}

// ListOption 查找选项
type ListOption struct {
	AuthorID   int64     // 作者ID
	LatestTime time.Time // 最新时间(分页)
	Limit      int       // 限制数量(分页)
	QueryAll   bool      // 是否查询所有(分页)
	CategoryID int64     // 分类ID
}

func (m *VideoModel) FindOneCategory(ctx context.Context, categoryID int64) (*Category, error) {
	var category Category
	err := m.db.WithContext(ctx).
		Where("id = ?", categoryID).
		First(&category).Error
	return &category, err
}

func (m *VideoModel) MustFindOneCategory(ctx context.Context, categoryID int64) *Category {
	ret, err := m.FindOneCategory(ctx, categoryID)
	if err != nil {
		panic(err)
	}
	return ret
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
		session = session.Where("created_at < ?", opts.LatestTime)
		session = session.Limit(opts.Limit)
	}
	return session.Order("created_at desc")
}

// List 通过时间点来获取比该时间点早的十个视频
func (m *VideoModel) List(ctx context.Context, opts ListOption) ([]*Video, error) {
	var result []*Video
	return result, m.applyOption(ctx, opts).Find(&result).Error
}

// Count 通过作者 ID 获取视频数量
func (m *VideoModel) Count(ctx context.Context, opts ListOption) (count int64, err error) {
	err = m.applyOption(ctx, opts).Count(&count).Error
	return
}

// ListVideoByAuthorId 通过作者 ID 获取视频列表
func (m *VideoModel) ListVideoByAuthorId(ctx context.Context, authorId uint) ([]*Video, error) {
	var result []*Video
	err := m.db.WithContext(ctx).Where("author_id = ?", authorId).Find(&result).Error
	return result, err
}

// ListVideoByCategoryId 通过分类 ID 获取视频列表
func (m *VideoModel) ListVideoByCategoryId(ctx context.Context, categoryId uint) ([]*Video, error) {
	var result []*Video
	err := m.db.WithContext(ctx).Where("category_id = ?", categoryId).Find(&result).Error
	return result, err
}

// FollowingUserVideo 关注用户的视频
func (m *VideoModel) FollowingUserVideo(ctx context.Context, doerId int64, opts ListOption) ([]*Video, error) {
	session := m.db.WithContext(ctx)
	session = session.
		Select("video.*").
		Joins("LEFT JOIN follow ON video.author_id = follow.followed_user_id ").
		Where("follow.user_id = ?", doerId)
	session = session.Where("`video`.created_at <= ?", opts.LatestTime)
	session = session.Limit(opts.Limit)
	var videos []*Video
	return videos, session.Find(&videos).Error
}

func (m *VideoModel) FindTagsByVideoId(ctx context.Context, videoId int64) ([]*Tag, error) {
	var tags []*Tag
	sess := m.db.WithContext(ctx)
	sess = sess.
		Select("`tags`.*").
		Joins("Join video_tags on `tags`.id = `video_tags`.tag_id").
		Where("`video_tags`.video_id = ?", videoId)
	return tags, sess.Find(&tags).Error
}

func (m *VideoModel) FindTagsByIds(ctx context.Context, ids []int64) ([]*Tag, error) {
	tags := []*Tag{} // 确保返回空数组
	err := m.db.WithContext(ctx).Model(&Tag{}).
		Where("`tags`.id IN ?", ids).
		Find(&tags).Error
	if err != nil {
		return tags, err
	}
	// 确保顺序
	var2idx := make(map[int64]int, len(ids))
	for i, v := range ids {
		var2idx[v] = i
	}
	sort.Slice(tags, func(i, j int) bool {
		return var2idx[tags[i].ID] < var2idx[tags[j].ID]
	})
	return tags, nil
}

func (m *VideoModel) MustFindTagsByVideoId(ctx context.Context, videoId int64) []*Tag {
	tags, err := m.FindTagsByVideoId(ctx, videoId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		panic(err)
	}
	return tags
}

func (m *VideoModel) IsAuthor(ctx context.Context, doerId int64, videoId int64) (bool, error) {
	var cnt int64
	err := m.db.WithContext(ctx).Model(&Video{}).
		Where("author_id = ? and id = ?", doerId, videoId).
		Count(&cnt).Error
	return cnt > 0, err
}

func (m *VideoModel) Delete(ctx context.Context, userId int64, videoId int64) error {
	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1.尝试删除视频
		rowsAffected := m.db.WithContext(ctx).
			Where("id = ? AND author_id = ?", videoId, userId).
			Delete(&Video{}).RowsAffected
		if rowsAffected == 0 {
			return fmt.Errorf("视频删除失败,请检查权限或视频是否存在")
		}
		// TODO: 清空点赞/评论/收藏 ?
		return nil
	})
}
