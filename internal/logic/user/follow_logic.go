package user

import (
	"context"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	doerUserId := jwt.GetUidFromCtx(l.ctx)
	res := l.svcCtx.UserModel.FollowUser(l.ctx, doerUserId, req.ContextUserId)
	data := fmt.Sprintf("Follow User ID: %d Success", req.ContextUserId)
	if errMq := l.svcCtx.KqPusherClient.Push(data); errMq != nil {
		logx.Errorf("FollowUser error: %v", errMq)
	}
	return nil, res
}
