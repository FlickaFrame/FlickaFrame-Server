package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
	"time"
)

type FollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowingLogic {
	return &FollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowingLogic) Following(req *types.FeedReq) (resp *types.FeedResp, err error) {
	doerId := jwt.GetUidFromCtx(l.ctx) // 从context中获取当前用户id
	LatestTime := time.Now()
	if req.Cursor != 0 {
		LatestTime = time.UnixMilli(req.Cursor)
	}
	videos, err := l.svcCtx.VideoModel.FollowingUserVideo(l.ctx, doerId, video_model.ListOption{
		AuthorID:   util.MustString2Int64(req.AuthorID),
		LatestTime: LatestTime,
		Limit:      req.Limit,
		QueryAll:   false,
		CategoryID: util.MustString2Int64(req.CategoryID),
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
	// 判断关注状态
	for i := range list {
		authorId, _ := strconv.ParseInt(list[i].VideoUserInfo.ID, 10, 64)
		resp.List[i].VideoUserInfo.IsFollow = l.svcCtx.UserModel.IsFollowing(l.ctx, doerId, authorId)
	}
	resp.IsEnd = len(resp.List) < req.Limit
	return
}
