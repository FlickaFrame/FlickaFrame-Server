package logic

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/oss"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/svc"
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
	endpoint, err := l.svcCtx.OssRpc.GetEndpoint(l.ctx, &oss.GetEndpointRequest{})
	if err != nil {
		return nil, err
	}
	resp = &types.OssEndpointResponse{
		EndPoint: endpoint.Endpoint,
	}
	return
}
