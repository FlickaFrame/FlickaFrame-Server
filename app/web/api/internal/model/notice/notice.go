package notice

import (
	"context"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/base"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"gorm.io/gorm"
	"time"
)

const (
	NoticeTypeFollow      = "follow"        // 关注消息
	NoticeTypeLikeVideo   = "like-video"    // 点赞消息(视频)
	NoticeTypeLikeComment = "like-comment"  // 点赞消息(评论)
	NoticeTypeComment     = "comment-video" // 评论消息
	NoticeCollect         = "collect"       // 收藏消息
	NoticeTypeReply       = "reply-comment" // 回复消息
	NoticeTypeAt          = "@-comment"     // @消息
	NoticeTypeSystem      = "system-notice" // 系统消息
)

const DefaultLimit = 10

type Notice struct {
	base.Model

	Checked    bool      `gorm:"default:0"` // 是否已读
	Content    string    // 内容
	NoticeType string    `gorm:"type:varchar(64)"` // 消息类型
	ToUserID   int64     `gorm:"index"`            // 接受用户id
	FromUserID int64     // 发送用户id
	NoticeTime time.Time // 通知时间
}

func (m *Notice) TableName() string {
	return "notices"
}

type NoticeModel struct {
	db *orm.DB
}

func NewNoticeModel(db *orm.DB) *NoticeModel {
	return &NoticeModel{
		db: db,
	}
}

func (m *NoticeModel) Insert(ctx context.Context, data *Notice) error {
	data.Model = base.NewModel()
	return m.db.WithContext(ctx).Create(data).Error
}

// ListOption 查找选项
type ListOption struct {
	AuthorID   int64     // 作者ID
	LatestTime time.Time // 最新时间(分页)
	Limit      int       // 限制数量(分页)
	QueryAll   bool      // 是否查询所有(分页)
	NoticeType string    // 通知类型
}

func (m *NoticeModel) applyOption(ctx context.Context, opts ListOption) *gorm.DB {
	session := m.db.WithContext(ctx)
	session = session.Where("to_user_id = ?", opts.AuthorID)
	// 根据通知类型
	if opts.NoticeType != "" {
		session = session.Where("notice_type = ?", opts.NoticeType)
	}
	// 分页
	if opts.Limit == 0 {
		opts.Limit = DefaultLimit
	}
	if !opts.QueryAll {
		session = session.Where("notice_time <= ?", opts.LatestTime)
		session = session.Limit(opts.Limit)
	}
	return session.Order("notice_time desc")
}

// List 通过时间点来获取比该时间点早的十个通知
func (m *NoticeModel) List(ctx context.Context, opts ListOption) ([]*Notice, error) {
	var result []*Notice
	return result, m.applyOption(ctx, opts).Find(&result).Error
}

type FollowNotice struct { // 关注消息
	FollowerId uint `json:"follower_id"` // 关注者id
}

func (n *FollowNotice) Message() string {
	return fmt.Sprintf("xxx关注了你")
}

type LikeNotice struct { // 点赞消息
	VideoId uint `json:"video_id"` // 视频id
}

func (n *LikeNotice) Message() string {
	return fmt.Sprintf("xxx点赞了你的xxx视频")
}

type CollectNotice struct {
	VideoId uint `json:"video_id"` // 视频id
	DoerId  uint `json:"doer_id"`  // 用户id
}

func (n *CollectNotice) Message() string {
	return fmt.Sprintf("xxx收藏了你的xxx视频")
}

type CommentNotice struct {
	VideoId uint `json:"video_id"` // 视频id
	DoerId  uint `json:"doer_id"`  // 用户id
}

func (n *CommentNotice) Message() string {
	return fmt.Sprintf("xxx评论了你的xxx视频")
}

type ReplyNotice struct {
	DoerId uint `json:"doer_id"` // 用户id
}

func (n *ReplyNotice) Message() string {
	return fmt.Sprintf("xxx回复了你的xxx评论")
}

type AtNotice struct {
	DoerId    uint `json:"doer_id"`    // 用户id
	CommentId uint `json:"comment_id"` // 评论id
}

func (n *AtNotice) Message() string {
	return fmt.Sprintf("xxx在xxx评论中@了你")
}

type SystemNotice struct {
}

func (n *SystemNotice) Message() string {
	return fmt.Sprintf("xxx")
}
