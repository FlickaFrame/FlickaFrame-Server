package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserDetailInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailInfoLogic {
	return &GetUserDetailInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserDetailInfoLogic) GetUserDetailInfo(req *types.UserDetailInfoReq) (resp *types.UserDetailInfoResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx)
	contextUserId := req.ContextUserId
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, contextUserId)
	if err != nil {
		return nil, err
	}
	basicInfo, err := NewConvert(l.ctx, l.svcCtx).BuildUserBasicInfo(l.ctx, user)
	if err != nil {
		return nil, err
	}
	interInfo, err := NewConvert(l.ctx, l.svcCtx).BuildUserInteractionInfo(l.ctx, doerId, contextUserId)
	if err != nil {
		return nil, err
	}
	statInfo, err := NewConvert(l.ctx, l.svcCtx).BuildUserStatisticalInfo(l.ctx, user)
	if err != nil {
		return nil, err
	}
	resp = &types.UserDetailInfoResp{
		UserBasicInfo:       *basicInfo,
		UserInteractionInfo: *interInfo,
		UserStatisticalInfo: *statInfo,
	}
	return
}
