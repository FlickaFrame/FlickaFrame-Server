package logic

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/pb/oss_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEndpointLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEndpointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEndpointLogic {
	return &GetEndpointLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEndpointLogic) GetEndpoint(in *oss_service.GetEndpointRequest) (*oss_service.GetEndpointResponse, error) {
	resp := &oss_service.GetEndpointResponse{
		Endpoint: l.svcCtx.Config.Oss.Endpoint,
	}
	return resp, nil
}
