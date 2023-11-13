package logic

import (
	"context"

	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowCountLogic {
	return &FollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowCountLogic) FollowCount(in *pb.FollowCountReq) (*pb.FollowCountResp, error) {
	ids, err := l.svcCtx.FollowCountModel.FindByUserIds(l.ctx, in.UserIds)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.FollowCountItem, 0, len(ids))
	for _, id := range ids {
		items = append(items, &pb.FollowCountItem{
			UserId:      id.UserID,
			FollowCount: int64(id.FollowCount),
			FansCount:   int64(id.FansCount),
		})
	}
	return &pb.FollowCountResp{
		Items: items,
	}, nil
}
