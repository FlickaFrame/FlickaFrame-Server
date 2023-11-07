package user

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/jinzhu/copier"

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
	doerId := jwt.GetUidFromCtx(l.ctx)
	doer := &user_model.User{}
	err = copier.Copy(&doer, req)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.UserModel.Update(l.ctx, doer, doerId)
	if err != nil {
		return nil, err
	}
	return NewCurrentUserInfoLogic(l.ctx, l.svcCtx).CurrentUserInfo(&types.UserDetailInfoReq{})
}
