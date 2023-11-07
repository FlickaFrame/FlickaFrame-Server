package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"time"

	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

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
	if req.PublishTime == 0 {
		req.PublishTime = time.Now().UnixMilli()
	}
	video := &video_model.Video{
		AuthorID:      doerId,
		Title:         req.Title,
		Description:   req.Description,
		PlayUrl:       req.PlayUrl,
		ThumbUrl:      req.ThumbUrl,
		PublishTime:   time.Unix(req.PublishTime/1000, req.PublishTime%1000), // TODO
		PublishStatus: publishStatus,
		Visibility:    req.Visibility,
		VideoDuration: req.VideoDuration,
		VideoHeight:   req.VideoHeight,
		VideoWidth:    req.VideoWidth,
		CategoryID:    util.MustString2Int64(req.CategoryID),
	}

	// 创建视频Tags
	for _, tag := range req.Tags {
		video.Tags = append(video.Tags, &video_model.Tag{
			Name: tag,
		})
	}
	err = l.svcCtx.VideoModel.Insert(l.ctx, video)
	return
}
