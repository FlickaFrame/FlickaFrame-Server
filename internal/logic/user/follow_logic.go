package user

import (
	"context"
	"time"
	"encoding/json"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	notice_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/notice"

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
	// data := fmt.Sprintf("Follow User ID: %d Success", req.ContextUserId)
	notice := notice_model.Notice{
		ToUserID:   req.ContextUserId,
		FromUserID: doerUserId,
		NoticeType: notice_model.NoticeTypeFollow,
		NoticeTime: time.Now(),
	}
	jsonBody, errJson := json.Marshal(notice)
	if errJson != nil {
		return nil, errJson
	}
	if errMq := l.svcCtx.KqPusherClient.Push(string(jsonBody)); errMq != nil {
		return nil, errMq
	}
	return nil, res
}
