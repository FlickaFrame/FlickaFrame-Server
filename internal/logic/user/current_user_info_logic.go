package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentUserInfoLogic {
	return &CurrentUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentUserInfoLogic) CurrentUserInfo() (resp *types.UserDetailInfoResp, err error) {
	userId := jwt.GetUidFromCtx(l.ctx)
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	return NewConvert(l.ctx, l.svcCtx).BuildUserDetailInfo(l.ctx, user, user)
}
