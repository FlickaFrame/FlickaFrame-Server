package comment

import (
	"context"
	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoCommentLogic {
	return &CreateVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVideoCommentLogic) CreateVideoComment(req *types.CreateVideoCommentReq) (resp *types.CreateVideoCommentResp, err error) {
	videoCount := video_count.NewVideoCountModel(l.svcCtx.BizRedis)
	resp = &types.CreateVideoCommentResp{}
	doer := jwt.GetUidFromCtx(l.ctx)
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
	comment, err := l.svcCtx.CommentModel.CreateParentComment(l.ctx, doer, util.MustString2Int64(req.VideoId), req.Content)
	if err != nil {
		return
	}
	parentComment, err := NewConvert(l.ctx, l.svcCtx).BuildParentComment(l.ctx, doer, comment)
	if err != nil {
		return nil, err
	}
	resp.Comment = parentComment
	return
}
