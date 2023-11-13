package favorite

import (
	"context"
	"encoding/json"
	notice_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/notice"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode/code"
	"github.com/zeromicro/go-zero/core/threading"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoLogic {
	return &FavoriteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteVideoLogic) FavoriteVideo(req *types.FavoriteReq) (resp *types.FavoriteResp, err error) {
	resp = &types.FavoriteResp{IsFavorite: true}
	doerId := jwt.GetUidFromCtx(l.ctx)
	// 检查视频是否存在
	video, err := l.svcCtx.VideoModel.FindOne(l.ctx, util.MustString2Int64(req.TargetId))
	if err != nil {
		logx.Info(err)
		return nil, code.VideoNotExistError
	}
	// redis incr
	count, err := l.svcCtx.BizRedis.Incr(cacheVideoLikeCount(util.MustString2Int64(req.TargetId)))
	resp.LikeCount = int(count)
	err = l.svcCtx.FavoriteModel.CreateVideoFavorite(l.ctx,
		doerId,
		util.MustString2Int64(req.TargetId))
	if err != nil {
		logx.Info(err)
		return nil, code.DuplicateFavoriteErr
	}

	threading.GoSafe(func() {
		notice := notice_model.Notice{
			ToUserID:   video.AuthorID,
			FromUserID: doerId,
			NoticeType: notice_model.NoticeTypeLikeVideo,
			NoticeTime: time.Now(),
		}
		jsonBody, errJson := json.Marshal(notice)
		if errJson != nil {
			logx.Error(errJson)
			return
		}
		if errMq := l.svcCtx.KqPusherClient.Push(string(jsonBody)); errMq != nil {
			logx.Error(errMq)
			return
		}
	})
	return resp, nil
}
