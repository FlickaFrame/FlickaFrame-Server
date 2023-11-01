package video

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVideoLogic {
	return &CreateVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVideoLogic) CreateVideo(req *types.CreateVideoReq) (resp *types.CreateVideoResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	publishStatus := 1 // TODO 更新发布状态
	video := &video_model.Video{
		AuthorID: doerId,
		CategoryID: req.CategoryID,
		Title: req.Title,
		Description: req.Description,
		PlayUrl: req.PlayUrl,
		ThumbUrl: req.ThumbUrl,
		PublishTime: req.PublishTime,
		PublishStatus: publishStatus,
		Visibility: req.Visibility,
	}
	err := l.svcCtx.VideoModel.Insert(l.ctx, video)
	return err
}
