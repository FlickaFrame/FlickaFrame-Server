package notice

import (
	"fmt"
	"time"
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/base"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
)

const (
	NoticeTypeFollow      = iota // 关注消息
	NoticeTypeLikeVideo          // 点赞消息(视频)
	NoticeTypeLikeComment        // 点赞消息(评论)
	NoticeTypeComment            // 评论消息
	NoticeCollect                // 收藏消息
	NoticeTypeReply              // 回复消息
	NoticeTypeAt                 // @消息
	NoticeTypeSystem             // 系统消息
)

type Notice struct {
	base.Model

	Checked    bool    `gorm:"default:0"`    // 是否已读
	Content    string                     	// 内容
	NoticeType int    											// 消息类型
	ToUserID   int64   `gorm:"index"`       // 接受用户id
	FromUserID int64      								  // 发送用户id
	NoticeTime time.Time 										// 通知时间
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
