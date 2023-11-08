package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoLogic) DeleteVideo(req *types.DeleteVideoReq) (resp *types.DeleteVideoResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	err = l.svcCtx.VideoModel.Delete(l.ctx, doerId, req.VideoID)
	return
}
