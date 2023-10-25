package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/code"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

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
	var userId uint
	userId, err = l.loginByPhone(req.Phone, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenerateToken(userId, l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		return nil, err
	}
	var tokenResp types.LoginResp
	err = copier.Copy(&tokenResp, token)
	return &tokenResp, err
}

func (l *LoginLogic) loginByPhone(mobile, password string) (uint, error) {
	user, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, mobile)
	if err != nil && !errors.Is(err, code.ErrNotFound) {
		return 0, errors.Wrapf(xcode.DB_ERROR, "根据手机号查询用户信息失败，mobile:%s,err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(code.ErrUserNoExistsError, "mobile:%s", mobile)
	}

	if !(util.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(code.ErrUsernamePwdError, "密码匹配出错")
	}

	return user.ID, nil
}
