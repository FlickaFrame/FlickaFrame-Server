package user

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListMyFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMyFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMyFollowingLogic {
	return &ListMyFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMyFollowingLogic) ListMyFollowing(req *types.ListMyFollowingReq) (resp *types.ListMyFollowingResp, err error) {
	userId := jwt.GetUidFromCtx(l.ctx)
	followings, err := l.svcCtx.UserModel.GetUserFollowing(l.ctx, userId, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  req.ListAll,
	})
	if err != nil {
		return nil, err
	}
	list, err := NewConvert(l.ctx, l.svcCtx).buildUserBasicInfoList(l.ctx, followings)
	resp = &types.ListMyFollowingResp{
		FollowUser: make([]*types.FollowUser, len(list)),
	}
	err = copier.Copy(&resp.FollowUser, &list)
	return
}
