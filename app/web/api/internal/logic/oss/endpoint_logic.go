package oss

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EndpointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndpointLogic {
	return &EndpointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EndpointLogic) Endpoint() (resp *types.OssEndpointResponse, err error) {
	resp = &types.OssEndpointResponse{
		EndPoint: l.svcCtx.Config.Oss.Endpoint,
	}
	return
}
