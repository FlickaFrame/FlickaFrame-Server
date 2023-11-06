package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/logic/user"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/comment"
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/comment"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
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
	targetComment, err := c.svcCtx.CommentModel.FindChildComment(ctx, id)
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
func (c *Convert) BuildChildComment(ctx context.Context, doerId int64, comment *comment.ChildComment) (resp *types.ChildComment, err error) {
	resp = &types.ChildComment{}
	err = copier.Copy(&resp.CommentBasicInfo, comment)
	resp.ID = strconv.FormatInt(comment.ID, 10)
	// 1.构造用户信息
	if comment.User == nil {
		comment.User, err = c.svcCtx.UserModel.FindOne(ctx, comment.UserID)
		if err != nil {
			logx.Info("find user fail: ", err)
			return nil, err
		}
	}
	resp.UserInfo, err = user.NewConvert(ctx, c.svcCtx).BuildUserBasicInfo(ctx, comment.User)
	//2. 构造回复评论信息
	if comment.ReplyID != 0 {
		resp.TargetComment = c.BuildTargetComment(ctx, comment.ReplyID)
	}
	// TODO: 构造点赞信息
	return
}

// BuildParentComment 用于构建视频评论
func (c *Convert) BuildParentComment(ctx context.Context, doerId int64, comment *comment.ParentComment) (resp *types.ParentComment, err error) {
	resp = &types.ParentComment{}
	// 1.构造评论基本信息
	err = copier.Copy(&resp.CommentBasicInfo, comment)
	if err != nil {
		return nil, err
	}
	resp.VideoID = strconv.FormatInt(comment.VideoID, 10)
	// 2. 构造一级评论的用户信息
	if comment.User == nil {
		comment.User, err = c.svcCtx.UserModel.FindOne(ctx, comment.UserID)
		if err != nil {
			logx.Info("find user fail: ", err)
			return nil, err
		}
	}
	resp.UserInfo, err = user.NewConvert(ctx, c.svcCtx).BuildUserBasicInfo(ctx, comment.User)
	// 3.构造二级评论信息
	err = comment.LoadChildComments(ctx, c.svcCtx.DB, comment_model.Option{})
	if err != nil {
		return nil, err
	}
	for _, childComment := range comment.ChildComments {
		childCommentInfo, err := c.BuildChildComment(ctx, doerId, childComment)
		if err != nil {
			return nil, err
		}
		resp.ChildComments = append(resp.ChildComments, childCommentInfo)
	}
	return
}

func (c *Convert) BuildParentCommentList(ctx context.Context, doerId int64, comments []*comment.ParentComment) (resp []*types.ParentComment, err error) {
	for _, parentComment := range comments {
		commentInfo, err := c.BuildParentComment(ctx, doerId, parentComment)
		if err != nil {
			return nil, err
		}
		resp = append(resp, commentInfo)
	}
	return
}
