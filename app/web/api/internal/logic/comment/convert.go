package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/logic/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/comment"
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/comment"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

// Convert 用于将model层的数据转换为logic层的数据
type Convert struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvert(ctx context.Context, svcCtx *svc.ServiceContext) *Convert {
	return &Convert{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// BuildTargetComment 回复评论时构造目标评论
func (c *Convert) BuildTargetComment(ctx context.Context, id int64) *types.TargetComment {
	resp := &types.TargetComment{}
	targetComment, err := c.svcCtx.CommentModel.FindOneComment(ctx, id)
	if err != nil {
		logx.Info("find comment fail: ", err)
		return nil
	}
	targetUser, err := c.svcCtx.UserModel.FindOne(ctx, targetComment.UserID)
	if err != nil {
		logx.Info("find user fail: ", err)
		return nil
	}
	resp.UserInfo, _ = user.NewConvert(ctx, c.svcCtx).BuildUserBasicInfo(ctx, targetUser)
	resp.ID = strconv.FormatInt(targetComment.ID, 10)
	return resp
}

// BuildChildComment 用于构建二级评论
func (c *Convert) BuildChildComment(ctx context.Context, doerId int64, comment *comment.Comment) (resp *types.ChildComment, err error) {
	resp = &types.ChildComment{}
	err = copier.Copy(&resp.CommentBasicInfo, comment)
	resp.ID = strconv.FormatInt(comment.ID, 10)
	// 1.构造用户信息
	author, err := c.svcCtx.UserModel.FindOne(ctx, comment.UserID)
	if err != nil {
		logx.Info("find user fail: ", err)
		return nil, err
	}
	resp.ShowTags = []string{}
	if ok, _ := c.svcCtx.VideoModel.IsAuthor(ctx, doerId, comment.VideoID); ok {
		resp.ShowTags = append(resp.ShowTags, comment_model.IsAuthor)
	}
	resp.UserInfo, err = user.NewConvert(ctx, c.svcCtx).BuildUserBasicInfo(ctx, author)
	//2. 构造回复评论信息
	if comment.ReplyID != 0 {
		resp.TargetComment = &types.TargetComment{}
		targetComment, err := c.svcCtx.CommentModel.FindOneComment(ctx, comment.ReplyID)
		userInfo, err := c.svcCtx.UserModel.FindOne(ctx, targetComment.UserID)
		if err != nil {
			return nil, err
		}
		resp.TargetComment.ID = strconv.FormatInt(comment.ReplyID, 10)
		resp.TargetComment.UserInfo, err = user.NewConvert(c.ctx, c.svcCtx).BuildUserBasicInfo(ctx, userInfo)
		if err != nil {
			return nil, err
		}
	}
	// 3.构造点赞信息统计
	count, err := c.svcCtx.FavoriteModel.CountByCommentId(ctx, comment.ID)
	if err != nil {
		return nil, err
	}
	comment.LikeCount = int(count)
	// 4.构造是否点赞信息
	resp.Liked, err = c.svcCtx.FavoriteModel.IsExist(ctx, doerId, comment.ID)
	return
}

// BuildChildCommentList 用于构建二级评论列表
func (c *Convert) BuildChildCommentList(ctx context.Context, doerId int64, comments []*comment.Comment) (resp []*types.ChildComment, err error) {
	resp = make([]*types.ChildComment, 0, len(comments))
	for i := range comments {
		var item *types.ChildComment
		item, err = c.BuildChildComment(ctx, doerId, comments[i])
		if err != nil {
			return
		}
		resp = append(resp, item)
	}
	return
}

// BuildParentComment 用于构建一级评论
func (c *Convert) BuildParentComment(ctx context.Context, doerId int64, comment *comment.Comment) (resp *types.ParentComment, err error) {
	resp = &types.ParentComment{}
	// 1.构造评论基本信息
	err = copier.Copy(&resp.CommentBasicInfo, comment)
	if err != nil {
		return nil, err
	}
	resp.VideoID = strconv.FormatInt(comment.VideoID, 10)
	resp.ShowTags = []string{}
	if ok, _ := c.svcCtx.VideoModel.IsAuthor(ctx, doerId, comment.VideoID); ok {
		resp.ShowTags = append(resp.ShowTags, comment_model.IsAuthor)
	}
	// 2. 构造一级评论的用户信息
	author, err := c.svcCtx.UserModel.FindOne(ctx, comment.UserID)
	if err != nil {
		logx.Info("find user fail: ", err)
		return nil, err
	}
	resp.UserInfo, err = user.NewConvert(ctx, c.svcCtx).BuildUserBasicInfo(ctx, author)
	// 3.构造二级评论信息
	resp.ChildCount, err = c.svcCtx.CommentModel.CountCommentByParentCommentId(ctx, comment.ID)
	if err != nil {
		return nil, err
	}
	opts := &comment_model.CommentOption{Paginator: &orm.ListOptions{}}
	opts.Limit = 3 // TODO : 默认显示3条子评论
	comment.ChildComments, err = c.svcCtx.CommentModel.FindChildCommentByCommentId(ctx,
		comment.ID,
		opts,
	)
	if err != nil {
		return nil, err
	}
	resp.ChildComments, err = c.BuildChildCommentList(ctx, doerId, comment.ChildComments)
	if err != nil {
		return nil, err
	}
	// 4.构造点赞信息
	count, err := c.svcCtx.FavoriteModel.CountByCommentId(ctx, comment.ID)
	if err != nil {
		return nil, err
	}
	resp.LikedCount = count
	// 5.构造是否点赞信息
	resp.Liked, err = c.svcCtx.FavoriteModel.IsExist(ctx, comment.ID, doerId)
	return
}

func (c *Convert) BuildParentCommentList(ctx context.Context, doerId int64, comments []*comment.Comment) (resp []*types.ParentComment, err error) {
	resp = make([]*types.ParentComment, 0, len(comments))
	for _, parentComment := range comments {
		commentInfo, err := c.BuildParentComment(ctx, doerId, parentComment)
		if err != nil {
			return resp, err
		}
		resp = append(resp, commentInfo)
	}
	return
}
