package logic

import (
	"context"
	"github.com/qiniu/go-sdk/v7/storage"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/pb/oss_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileAccessUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileAccessUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileAccessUrlLogic {
	return &GetFileAccessUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileAccessUrlLogic) GetFileAccessUrl(in *oss_service.GetFileAccessUrlRequest) (*oss_service.GetFileAccessUrlResponse, error) {
	resp := &oss_service.GetFileAccessUrlResponse{}
	resp.Url = storage.MakePublicURL(l.svcCtx.Config.Oss.Endpoint, in.Key)
	return resp, nil
}
