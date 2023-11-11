package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearPlayHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearPlayHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearPlayHistoryLogic {
	return &ClearPlayHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ClearPlayHistory 清空播放历史
func (l *ClearPlayHistoryLogic) ClearPlayHistory() (resp *types.ClearPlayHistoryResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.VideoHistoryModel.ClearPlayHistory(l.ctx, doerId)
	if err != nil {
		return nil, err
	}
	return
}
