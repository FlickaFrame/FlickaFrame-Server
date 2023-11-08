package logic

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/pb/user_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListByIdsLogic {
	return &ListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListByIdsLogic) ListByIds(in *user_service.ListByIdsRequest) (*user_service.ListByIdsResponse, error) {
	// todo: add your logic here and delete this line

	return &user_service.ListByIdsResponse{}, nil
}
