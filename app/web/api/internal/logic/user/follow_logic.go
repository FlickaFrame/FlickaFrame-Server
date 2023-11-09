package user

import (
	"context"
	"encoding/json"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"time"

	notice_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/notice"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

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
	_, err = l.svcCtx.FollowRpc.Follow(l.ctx, &follow_rpc.FollowRequest{
		UserId:         doerUserId,
		FollowedUserId: req.ContextUserId,
	})
	if err != nil {
		return nil, err
	}

	notice := notice_model.Notice{
		ToUserID:   req.ContextUserId,
		FromUserID: doerUserId,
		NoticeType: notice_model.NoticeTypeFollow,
		NoticeTime: time.Now(),
	}
	jsonBody, err := json.Marshal(notice)
	if err != nil {
		return nil, err
	}
	if errMq := l.svcCtx.KqPusherClient.Push(string(jsonBody)); errMq != nil {
		return nil, errMq
	}
	return
}
