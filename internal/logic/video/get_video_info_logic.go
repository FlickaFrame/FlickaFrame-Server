package video

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoInfoLogic {
	return &GetVideoInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoInfoLogic) GetVideoInfo(req *types.GetVideoInfoReq) (resp *types.GetVideoInfoResp, err error) {
	video, err := l.svcCtx.VideoModel.FindOne(l.ctx, req.VideoId)
	if err != nil {
		logx.Info(err)
		return nil, err
	}
	videoBasicInfo, err := NewConvert(l.ctx, l.svcCtx).BuildVideoBasicInfo(l.ctx, video)
	resp = &types.GetVideoInfoResp{
		Video: videoBasicInfo,
	}
	return
}
