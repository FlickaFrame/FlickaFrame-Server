package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	//doerId := jwt.GetUidFromCtx(l.ctx) // 从context中获取当前用户id
	LatestTime := time.Now()
	if req.Cursor != 0 {
		LatestTime = time.UnixMilli(req.Cursor)
	}
	videos, err := l.svcCtx.VideoModel.List(l.ctx, video_model.ListOption{
		AuthorID:   req.AuthorID,
		LatestTime: LatestTime,
		Limit:      req.Limit,
		QueryAll:   false,
		CategoryID: req.CategoryID,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("find videos by latest time error: %v", err)
		return
	}
	list, err := NewConvert(l.ctx, l.svcCtx).BuildVideoBasicInfoList(l.ctx, videos)
	if err != nil {
		return nil, err
	}

	// 判断是否无视频
	nextTime := LatestTime.UnixMilli()
	if len(videos) > 0 {
		nextTime = videos[len(videos)-1].CreatedAt.UnixMilli() - 1
	}
	resp = &types.FeedResp{
		List: make([]*types.FeedVideoItem, len(list)),
		Next: strconv.FormatInt(nextTime, 10),
	}
	err = copier.Copy(&resp.List, &list)
	return
}
