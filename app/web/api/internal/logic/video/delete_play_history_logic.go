package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePlayHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePlayHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePlayHistoryLogic {
	return &DeletePlayHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePlayHistoryLogic) DeletePlayHistory(req *types.DeletePlayHistoryReq) (resp *types.PlayHistoryResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.VideoHistoryModel.DeleteSpecificPlayHistory(l.ctx, doerId, videoId)
	if err != nil {
		return nil, err
	}
	return
}
