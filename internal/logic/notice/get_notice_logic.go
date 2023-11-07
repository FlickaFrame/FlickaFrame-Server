package notice

import (
	"context"
	"time"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	notice_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/notice"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/jinzhu/copier"
	"strconv"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeLogic {
	return &GetNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNoticeLogic) GetNotice(req *types.FollowNoticeReq) (resp *types.FollowNoticeResp, err error) {
	doerUserId := jwt.GetUidFromCtx(l.ctx)
	LatestTime := time.Now()
	if req.Cursor != 0 {
		LatestTime = time.UnixMilli(req.Cursor)
	}
	notices, err := l.svcCtx.NoticeModel.List(l.ctx, notice_model.ListOption{
		AuthorID:   doerUserId,
		LatestTime: LatestTime,
		Limit:      req.Limit,
		QueryAll:   false,
		NoticeType: req.NoticeType,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("find notices by latest time error: %v", err)
		return
	}
	nextTime := LatestTime.UnixMilli()
	if len(notices) > 0 {
		nextTime = notices[len(notices)-1].NoticeTime.UnixMilli() - 1
	}
	resp = &types.FollowNoticeResp{
		List: make([]*types.NoticeItem, len(notices)),
		Next: strconv.FormatInt(nextTime, 10),
	}
	err = copier.Copy(&resp.List, &notices)
	for i := 0; i < len(notices); i++ {
		resp.List[i].NoticeTime = notices[i].NoticeTime.UnixMilli()
		// TODO
	}
	resp.IsEnd = len(resp.List) < req.Limit
	return
}
