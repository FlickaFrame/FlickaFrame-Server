// Code generated by goctl. DO NOT EDIT.
// Source: oss.proto

package server

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/logic"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/pb/oss_service"
)

type OssServer struct {
	svcCtx *svc.ServiceContext
	oss_service.UnimplementedOssServer
}

func NewOssServer(svcCtx *svc.ServiceContext) *OssServer {
	return &OssServer{
		svcCtx: svcCtx,
	}
}

func (s *OssServer) CreatUpToken(ctx context.Context, in *oss_service.CreateUpTokenRequest) (*oss_service.CreateUpTokenResponse, error) {
	l := logic.NewCreatUpTokenLogic(ctx, s.svcCtx)
	return l.CreatUpToken(in)
}

func (s *OssServer) GetEndpoint(ctx context.Context, in *oss_service.GetEndpointRequest) (*oss_service.GetEndpointResponse, error) {
	l := logic.NewGetEndpointLogic(ctx, s.svcCtx)
	return l.GetEndpoint(in)
}

func (s *OssServer) GetFileAccessUrl(ctx context.Context, in *oss_service.GetFileAccessUrlRequest) (*oss_service.GetFileAccessUrlResponse, error) {
	l := logic.NewGetFileAccessUrlLogic(ctx, s.svcCtx)
	return l.GetFileAccessUrl(in)
}