package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChildCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateChildCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChildCommentLogic {
	return &CreateChildCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateChildCommentLogic) CreateChildComment(req *types.CreateChildCommentReq) (resp *types.CreateChildCommentResp, err error) {
	resp = &types.CreateChildCommentResp{}
	doer := jwt.GetUidFromCtx(l.ctx)
	// 检测视频是否存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, req.VideoId)
	if err != nil {
		logx.Info(err)
		return nil, code.VideoNotExistError
	}
	comment, err := l.svcCtx.CommentModel.CreateChildComment(l.ctx, doer,
		req.VideoId, req.Content, req.ParentCommentId, req.TargetCommentId)
	if err != nil {
		return
	}
	parentComment, err := NewConvert(l.ctx, l.svcCtx).BuildChildComment(l.ctx, doer, comment)
	if err != nil {
		return nil, err
	}
	resp.Comment = parentComment
	return
}
