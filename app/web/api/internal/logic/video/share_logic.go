package video

import (
	"context"
	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareLogic {
	return &ShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareLogic) Share(req *types.ShareVideoReq) (resp *types.ShareVideoResp, err error) {
	// 确保视频存在
	_, err = l.svcCtx.VideoModel.FindOne(l.ctx, req.VideoId)
	if err != nil {
		return nil, code.VideoNotExistError
	}
	count, err := video_count.
		NewVideoCountModel(l.svcCtx.BizRedis).
		IncrShareCount(l.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	resp = &types.ShareVideoResp{ShareCount: count}
	return
}
