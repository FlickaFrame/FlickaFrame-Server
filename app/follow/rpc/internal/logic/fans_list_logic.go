package logic

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/model"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
	"github.com/zeromicro/go-zero/core/threading"
	"strconv"
	"time"

	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

const userFansExpireTime = 3600 * 24 * 2

type FansListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FansListLogic {
	return &FansListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FansList 获取粉丝列表
func (l *FansListLogic) FansList(in *pb.FansListRequest) (*pb.FansListResponse, error) {
	// 1. 参数校验
	if in.FollowedUserId == 0 {
		return nil, code.UserIdEmpty
	}
	if in.PageSize == 0 {
		in.PageSize = types.DefaultPageSize
	}
	if in.Cursor == 0 {
		in.Cursor = time.Now().UnixMilli()
	}

	var (
		err            error
		isCache, isEnd bool
		lastId, cursor int64
		fansUserIds    []int64
		follows        []*model.Follow
		curPage        []*pb.FansItem
	)

	userIds, _ := l.getFansUserIdsFromCache(l.ctx, in.FollowedUserId, in.Cursor, in.PageSize)
	if len(userIds) > 0 {
		isCache = true
		if userIds[len(userIds)-1] == -1 { // 剔除分页占位符
			userIds = userIds[:len(userIds)-1]
			isEnd = true
		}
		if len(userIds) == 0 {
			return &pb.FansListResponse{}, nil
		}
		follows, err = l.svcCtx.FollowModel.FindByFollowedUserIds(l.ctx, userIds)
		if err != nil {
			l.Logger.Errorf("[FansList] FollowModel.FindByFollowedUserIds error: %v req: %v", err, in)
			return nil, err
		}
		for _, follow := range follows {
			userIds = append(userIds, follow.UserID)
			curPage = append(curPage, &pb.FansItem{
				Id:         follow.ID,
				UserId:     follow.UserID,
				CreateTime: follow.CreatedAt.UnixMilli(),
			})
		}
	} else { // 缓存不存在 => 从数据库中获取(并且刷新缓存分页)
		follows, err = l.svcCtx.FollowModel.FindByFollowedUserId(l.ctx, in.FollowedUserId, types.CacheMaxFansCount)
		if err != nil {
			l.Logger.Errorf("[FansList] FollowModel.FindByFollowedUserId error: %v req: %v", err, in)
			return nil, err
		}
		if len(follows) == 0 {
			return &pb.FansListResponse{}, nil
		}
		var firstPageFollows []*model.Follow
		if len(follows) > int(in.PageSize) {
			firstPageFollows = follows[:in.PageSize]
		} else {
			firstPageFollows = follows
			isEnd = true
		}
		for _, follow := range firstPageFollows {
			fansUserIds = append(fansUserIds, follow.UserID)
			curPage = append(curPage, &pb.FansItem{
				Id:         follow.ID,
				UserId:     follow.UserID,
				CreateTime: follow.CreatedAt.UnixMilli(),
			})
		}
	}
	if len(curPage) > 0 {
		pageLast := curPage[len(curPage)-1]
		lastId = pageLast.Id
		cursor = pageLast.CreateTime
		if cursor < 0 {
			cursor = 0
		}
		for k, follow := range curPage {
			if follow.CreateTime == in.Cursor && follow.Id == in.Id {
				curPage = curPage[k:]
				break
			}
		}
	}
	fc, err := l.svcCtx.FollowCountModel.FindByUserIds(l.ctx, fansUserIds) // 获取粉丝数
	if err != nil {
		l.Logger.Errorf("[FollowList] FollowCountModel.FindByUserIds error: %v followedUserIds: %v", err, fansUserIds)
	}
	uidFansCount := make(map[int64]int)
	for _, f := range fc {
		uidFansCount[f.UserID] = f.FansCount
	}
	for _, cur := range curPage {
		cur.FansUserId = int64(uidFansCount[cur.FansUserId])
	}
	ret := &pb.FansListResponse{
		IsEnd:  isEnd,
		Cursor: cursor,
		Id:     lastId,
		Items:  curPage,
	}

	if !isCache {
		threading.GoSafe(func() {
			if len(follows) < types.CacheMaxFansCount && len(follows) > 0 {
				follows = append(follows, &model.Follow{UserID: -1}) // 分页占位符
			}
			err = l.addCacheFans(context.Background(), in.FollowedUserId, follows) // 异步刷新缓存
			if err != nil {
				logx.Errorf("addCacheFollow error: %v", err)
			}
		})
	}

	return ret, nil
}

// getFansUserIdsFromCache 获取缓存中的用户id列表
func (l *FansListLogic) getFansUserIdsFromCache(ctx context.Context, userId, cursor, pageSize int64) ([]int64, error) {
	key := userFansKey(userId)
	exist, err := l.svcCtx.BizRedis.ExistsCtx(ctx, key)
	if err != nil {
		logx.Errorf("[getFansUserIdsFromCache] BizRedis.ExistsCtx error: %v", err)
	}
	if exist { // 缓存存在
		err = l.svcCtx.BizRedis.ExpireCtx(ctx, key, userFansExpireTime)
		if err != nil {
			logx.Errorf("[getFansUserIdsFromCache] BizRedis.ExpireCtx error: %v", err)
		}
	}
	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, key, 0, cursor, 0, int(pageSize))
	if err != nil {
		logx.Errorf("[getFansUserIdsFromCache] BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx error: %v", err)
		return nil, err
	}
	var uids []int64
	for _, pair := range pairs {
		uid, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			logx.Errorf("[getFansUserIdsFromCache] strconv.ParseInt error: %v", err)
			continue
		}
		uids = append(uids, uid)
	}

	return uids, nil
}

// addCacheFans 将粉丝关系写入缓存
func (l *FansListLogic) addCacheFans(ctx context.Context, userId int64, follows []*model.Follow) error {
	if len(follows) == 0 {
		return nil
	}
	key := userFansKey(userId)
	for _, follow := range follows {
		var score int64
		if follow.UserID == -1 {
			score = 0
		} else {
			score = follow.CreatedAt.UnixMilli()
		}
		_, err := l.svcCtx.BizRedis.ZaddCtx(ctx, key, score, strconv.FormatInt(follow.UserID, 10))
		if err != nil {
			logx.Errorf("[addCacheFans] BizRedis.ZaddCtx error: %v", err)
			return err
		}
	}

	return l.svcCtx.BizRedis.ExpireCtx(ctx, key, userFansExpireTime) // 刷新缓存过期时间
}
