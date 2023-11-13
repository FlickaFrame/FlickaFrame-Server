package user

import (
	"context"
	user_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RankingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRankingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RankingLogic {
	return &RankingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Ranking User Ranking By Score(ViewTimes, LikeTimes, CommentTimes,UpdateFrequency)
func (l *RankingLogic) Ranking(req *types.RankingReq) (resp *types.RankingResp, err error) {
	//TODO: implement ranking by score(ViewTimes, LikeTimes, CommentTimes,UpdateFrequency)
	users, err := l.svcCtx.UserModel.List(l.ctx, user_model.ListOptions{
		ListOptions: orm.ListOptions{
			PageSize: req.PageSize,
			Page:     req.Page,
			ListAll:  req.ListAll,
		},
		UserIds: nil,
	})
	resp = &types.RankingResp{
		Users: make([]*types.UserBasicInfo, 0, len(users)),
	}
	err = copier.Copy(&resp.Users, &users)
	return
}
