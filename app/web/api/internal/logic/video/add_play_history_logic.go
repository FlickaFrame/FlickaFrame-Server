package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPlayHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPlayHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPlayHistoryLogic {
	return &AddPlayHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddPlayHistory 添加播放历史
func (l *AddPlayHistoryLogic) AddPlayHistory(req *types.PlayHistoryReq) (resp *types.PlayHistoryResp, err error) {
	resp = &types.PlayHistoryResp{}
	doerId := jwt.GetUidFromCtx(l.ctx)
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.VideoHistoryModel.AddCachePlayHistory(l.ctx, doerId, videoId)
	return
}
