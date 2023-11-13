package comment

import (
	"context"
	comment_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/comment"
	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/jinzhu/copier"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListVideoCommentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListVideoCommentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVideoCommentsLogic {
	return &ListVideoCommentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListVideoCommentsLogic) ListVideoComments(req *types.ListVideoCommentsReq) (resp *types.ListVideoCommentsResp, err error) {
	videoCount := video_count.NewVideoCountModel(l.svcCtx.BizRedis)
	resp = &types.ListVideoCommentsResp{}
	doerId := jwt.GetUidFromCtx(l.ctx)
	opts := &comment_model.CommentOption{Paginator: &orm.ListOptions{}}
	err = copier.Copy(&opts, req)
	if err != nil {
		return nil, err
	}
	comments, err := l.svcCtx.CommentModel.FindParentCommentByVideoId(l.ctx, util.MustString2Int64(req.VideoId), opts)
	resp.Comments, err = NewConvert(l.ctx, l.svcCtx).BuildParentCommentList(l.ctx, doerId, comments)
	videoInt, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}
	resp.Total, err = videoCount.GetCommentCount(l.ctx, videoInt)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
