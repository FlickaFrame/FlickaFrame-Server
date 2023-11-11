package video

import (
	"context"
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/jinzhu/copier"
	"strconv"
	"time"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlayHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPlayHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlayHistoryLogic {
	return &GetPlayHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPlayHistoryLogic) GetPlayHistory(req *types.FeedReq) (resp *types.FeedResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx) // 从context中获取当前用户id
	resp = &types.FeedResp{List: make([]*types.FeedVideoItem, 0)}
	if req.Cursor == 0 {
		req.Cursor = time.Now().UnixMilli()
	}
	histories, err := l.svcCtx.VideoHistoryModel.GetPlayHistoryFromCache(l.ctx, doerId, req.Cursor, int64(req.Limit))
	if err != nil {
		return nil, err
	}
	videoIdx := make([]int64, 0, len(histories))
	for _, history := range histories {
		videoIdx = append(videoIdx, history.VideoId)
	}
	if len(videoIdx) > 0 && videoIdx[len(videoIdx)-1] == -1 {
		histories = histories[:len(histories)-1]
		videoIdx = videoIdx[:len(videoIdx)-1]
		resp.IsEnd = true
	}
	videos, err := l.svcCtx.VideoModel.FindByIDs(l.ctx, videoIdx)
	if err != nil {
		return nil, err
	}
	if err != nil {
		logx.WithContext(l.ctx).Errorf("find videos by latest time error: %v", err)
		return
	}
	list, err := NewConvert(l.ctx, l.svcCtx).BuildVideoBasicInfoList(l.ctx, videos)
	if err != nil {
		return nil, err
	}

	// 判断是否无视频
	resp.Next = strconv.FormatInt(time.Now().UnixMilli(), 10)
	if len(histories) > 0 {
		resp.Next = strconv.FormatInt(histories[len(histories)-1].PlayTime.UnixMilli(), 10)
	}
	err = copier.Copy(&resp.List, &list)
	// 判断关注状态
	for i := range list {
		authorId, _ := strconv.ParseInt(list[i].VideoUserInfo.ID, 10, 64)
		follow, err := l.svcCtx.FollowRpc.IsFollow(l.ctx, &follow_rpc.IsFollowReq{
			UserId:         doerId,
			FollowedUserId: authorId,
		})
		if err != nil {
			return nil, err
		}
		resp.List[i].VideoUserInfo.IsFollow = follow.IsFollow
	}
	return
}
