package user

import (
	"context"
	"encoding/json"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	user_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
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
	// 判断关注用户是否存在
	user, err := l.svcCtx.UserRpc.FindById(l.ctx, &user_rpc.FindByIdRequest{
		UserId: req.ContextUserId,
	})
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, code.ErrUserNoExistsError
	}
	// 关注操作
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
