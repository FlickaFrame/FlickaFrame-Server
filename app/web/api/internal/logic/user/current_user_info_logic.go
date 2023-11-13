package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

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

// CurrentUserInfo 获取当前用户信息
func (l *CurrentUserInfoLogic) CurrentUserInfo(req *types.UserDetailInfoReq) (resp *types.UserDetailInfoResp, err error) {
	req.ContextUserId = jwt.GetUidFromCtx(l.ctx)
	return NewGetUserDetailInfoLogic(l.ctx, l.svcCtx).GetUserDetailInfo(req)
}
