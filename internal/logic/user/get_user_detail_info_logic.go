package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailInfoLogic {
	return &GetUserDetailInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailInfoLogic) GetUserDetailInfo(req *types.UserDetailInfoReq) (resp *types.UserDetailInfoResp, err error) {
	userId := req.ContextUserId
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	return NewConvert(l.ctx, l.svcCtx).BuildUserDetailInfo(l.ctx, user, user)
}
