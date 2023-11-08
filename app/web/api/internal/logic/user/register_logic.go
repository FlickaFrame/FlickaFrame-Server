package user

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
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
	// TODO: Use DB Unique index to ensure that the phone number is unique
	user, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, req.Phone)
	if err != nil && !errors.Is(err, code.ErrNotFound) {
		return nil, errors.Wrapf(xcode.DB_ERROR, "mobile:%s,err:%v", req.Phone, err)
	}
	if user.ID != 0 {
		return nil, errors.Wrapf(code.ErrUserAlreadyRegisterError, "Register user exists mobile:%s,err:%v", req.Phone, err)
	}
	user = &user_model.User{
		NickName: req.NickName,
		Phone:    req.Phone,
		Password: req.Password,
	}
	if len(user.NickName) == 0 { // If the nickname is empty, generate a random nickname
		user.NickName = util.KRand(8, util.KC_RAND_KIND_ALL)
	}
	user.Password = util.Md5ByString(user.Password)
	err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, errors.Wrapf(xcode.DB_ERROR, "Register db user Insert err:%v,user:%+v", err, user)
	}

	//2、Generate the token, so that the service doesn't call rpc internally
	token, err := jwt.GenerateToken(user.ID, l.svcCtx.Config.JwtAuth.AccessSecret, l.svcCtx.Config.JwtAuth.AccessExpire)
	if err != nil {
		return nil, err
	}
	var tokenResp types.RegisterResp
	err = copier.Copy(&tokenResp, token)
	return &tokenResp, err
}
