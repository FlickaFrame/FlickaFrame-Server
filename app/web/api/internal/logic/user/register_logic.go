package user

import (
	"context"
	user_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	userInfo, err := l.svcCtx.UserRpc.FindByMobile(l.ctx, &user_rpc.FindByMobileRequest{
		Mobile: req.Phone,
	})
	if err != nil { // 数据库查询失败
		return nil, errors.Wrapf(xcode.DB_ERROR, "mobile:%s,err:%v", req.Phone, err)
	}
	if userInfo.Id != 0 { // 用户已存在
		return nil, errors.Wrapf(code.ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", req.Phone, err)
	}
	user := &user_rpc.RegisterRequest{
		Nickname: req.NickName,
		Phone:    req.Phone,
		Avatar:   "",
		Password: req.Password,
	}
	registerRsp, err := l.svcCtx.UserRpc.Register(l.ctx, user)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, errors.Wrapf(xcode.DB_ERROR, "Register db user Insert err:%v,user:%+v", err, user)
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	token, err := jwt.GenerateToken(registerRsp.UserId, l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		return nil, err
	}
	var tokenResp types.RegisterResp
	err = copier.Copy(&tokenResp, token)
	return &tokenResp, err
}
