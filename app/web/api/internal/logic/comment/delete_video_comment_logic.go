package comment

import (
	"context"
	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoCommentLogic {
	return &DeleteVideoCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoCommentLogic) DeleteVideoComment(req *types.DeleteVideoCommentReq) (resp *types.DeleteVideoCommentResp, err error) {
	videoCount := video_count.NewVideoCountModel(l.svcCtx.BizRedis)

	doerId := jwt.GetUidFromCtx(l.ctx)
	videoInt, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}
	count, err := videoCount.DecrCommentCount(l.ctx, videoInt)
	if err != nil {
		return nil, err
	}
	_ = count // TODO: 返回播放量给前端
	// TODO:评论数量-1
	err = l.svcCtx.CommentModel.DeleteCommentByDoer(l.ctx, util.MustString2Int64(req.CommentId), doerId)
	return nil, err
}
