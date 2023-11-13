package comment

import (
	"context"
	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

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
	videoCount := video_count.NewVideoCountModel(l.svcCtx.BizRedis)
	resp = &types.CreateChildCommentResp{}
	doer := jwt.GetUidFromCtx(l.ctx)
	// 检测视频是否存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, util.MustString2Int64(req.VideoId))
	if err != nil {
		logx.Info(err)
		return nil, code.VideoNotExistError
	}
	videoInt, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}
	count, err := videoCount.IncrCommentCount(l.ctx, videoInt)
	if err != nil {
		return nil, err
	}
	_ = count // TODO 返回评论数给前端
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
