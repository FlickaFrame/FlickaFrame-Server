package logic

import (
	"context"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
	"github.com/pkg/errors"

	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/pb/user_service"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user_service.LoginRequest) (*user_service.LoginResponse, error) {
	resp := &user_service.LoginResponse{}
	if in.Phone == "" || in.Password == "" {
		return resp, fmt.Errorf("手机号或密码不能为空")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		return resp, err
	}
	if !(util.Md5ByString(in.Password) == userInfo.Password) {
		return nil, errors.Wrap(code.ErrUsernamePwdError, "密码匹配出错")
	}
	resp.UserId = userInfo.ID
	return resp, nil
}
