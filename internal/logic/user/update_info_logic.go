package user

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"time"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateInfoReq) (resp *types.UserDetailInfoResp, err error) {
	doerUserId := jwt.GetUidFromCtx(l.ctx)
	doer := &user_model.User{
		ID:        doerUserId,
		UpdatedAt: time.Now(),
		NickName:  req.NickName,
		Age:       req.Age,
		Gender:    req.Gender,
		Slogan:    req.Slogan,
	}
	err = l.svcCtx.UserModel.Update(l.ctx, doer)
	if err != nil {
		return nil, err
	}
	return NewCurrentUserInfoLogic(l.ctx, l.svcCtx).CurrentUserInfo(&types.UserDetailInfoReq{})
}
