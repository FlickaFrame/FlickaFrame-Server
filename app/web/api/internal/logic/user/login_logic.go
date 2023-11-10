package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = &types.LoginResp{}
	userId, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil || userId.UserId == 0 {
		return nil, err
	}
	token, err := jwt.GenerateToken(userId.UserId, l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		return nil, err
	}
	var tokenResp types.LoginResp
	err = copier.Copy(&tokenResp, token)
	return &tokenResp, err
}
