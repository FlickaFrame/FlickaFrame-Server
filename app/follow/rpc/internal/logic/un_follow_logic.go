package logic

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/model"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
	"gorm.io/gorm"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnFollowLogic) UnFollow(in *pb.UnFollowRequest) (*pb.UnFollowResponse, error) {
	// 1. 参数校验
	if in.UserId == 0 {
		return nil, code.FollowUserIdEmpty
	}
	if in.FollowedUserId == 0 {
		return nil, code.FollowedUserIdEmpty
	}
	// 2. 数据库操作
	follow, err := l.svcCtx.FollowModel.FindByUserIDAndFollowedUserID(l.ctx, in.UserId, in.FollowedUserId)
	if err != nil {
		l.Logger.Errorf("[UnFollow] FollowModel.FindByUserIDAndFollowedUserID err: %v req: %v", err, in)
		return nil, err
	}
	if follow == nil {
		return nil, nil
	}
	if follow.FollowStatus == types.FollowStatusUnfollow {
		return nil, nil
	}

	// 事务
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		err := model.NewFollowModel(tx).
			UpdateFields(l.ctx, follow.ID, map[string]interface{}{
				"follow_status": types.FollowStatusUnfollow,
			})
		if err != nil {
			return err
		}
		err = model.NewFollowCountModel(tx).
			DecrFollowCount(l.ctx, in.UserId)
		if err != nil {
			return err
		}
		return model.NewFollowCountModel(tx).
			DecrFansCount(l.ctx, in.FollowedUserId)
	})
	if err != nil {
		l.Logger.Errorf("[UnFollow] Transaction error: %v", err)
		return nil, err
	}
	// 3. 缓存操作
	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, userFollowKey(in.UserId), strconv.FormatInt(in.FollowedUserId, 10))
	if err != nil {
		l.Logger.Errorf("[UnFollow] BizRedis.ZremCtx error: %v", err)
		return nil, err
	}
	_, err = l.svcCtx.BizRedis.ZremCtx(l.ctx, userFansKey(in.FollowedUserId), strconv.FormatInt(in.UserId, 10))
	if err != nil {
		l.Logger.Errorf("[UnFollow] BizRedis.ZremCtx error: %v", err)
		return nil, err
	}

	return &pb.UnFollowResponse{}, nil
}
