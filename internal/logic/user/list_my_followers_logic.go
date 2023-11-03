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

type ListMyFollowersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListMyFollowersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListMyFollowersLogic {
	return &ListMyFollowersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListMyFollowersLogic) ListMyFollowers(req *types.ListMyFollowersReq) (resp *types.ListMyFollowersResp, err error) {
	userID := jwt.GetUidFromCtx(l.ctx)
	followers, err := l.svcCtx.UserModel.GetUserFollowers(l.ctx, userID, orm.ListOptions{
		PageSize: req.PageSize,
		Page:     req.Page,
		ListAll:  req.ListAll,
	})
	if err != nil {
		return nil, err
	}
	list, err := NewConvert(l.ctx, l.svcCtx).buildUserBasicInfoList(l.ctx, followers)
	resp = &types.ListMyFollowersResp{
		FollowUser: make([]*types.FollowUser, len(list)),
	}
	err = copier.Copy(&resp.FollowUser, &list)
	return
}
