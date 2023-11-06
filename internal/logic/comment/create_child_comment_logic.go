package comment

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"strconv"

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
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, util.MustString2Int64(req.VideoId))
	if err != nil {
		logx.Info(err)
		return nil, code.VideoNotExistError
	}
	ParentCommentId, _ := strconv.ParseInt(req.ParentCommentId, 10, 64)
	TargetCommentId, _ := strconv.ParseInt(req.TargetCommentId, 10, 64)

	comment, err := l.svcCtx.CommentModel.CreateChildComment(l.ctx, doer,
		util.MustString2Int64(req.VideoId), req.Content, ParentCommentId, TargetCommentId)
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
