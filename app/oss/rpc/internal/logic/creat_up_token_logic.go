package logic

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/pb/oss_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatUpTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatUpTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatUpTokenLogic {
	return &CreatUpTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatUpTokenLogic) CreatUpToken(in *oss_service.CreateUpTokenRequest) (*oss_service.CreateUpTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &oss_service.CreateUpTokenResponse{}, nil
}
