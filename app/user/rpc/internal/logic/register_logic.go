package logic

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/model"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/pb/user_service"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user_service.RegisterRequest) (*user_service.RegisterResponse, error) {
	var user = &model.User{
		NickName:  in.Nickname,
		AvatarUrl: in.Avatar,
		Password:  in.Password,
		Phone:     in.Phone,
	}
	if len(user.NickName) == 0 { // If the nickname is empty, generate a random nickname
		user.NickName = util.KRand(8, util.KC_RAND_KIND_ALL)
	}
	if len(user.Password) > 0 {
		user.Password = util.Md5ByString(user.Password)
	}
	err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}
	return &user_service.RegisterResponse{
		UserId: user.ID,
	}, nil
}
