package video

import (
	"context"
	"strconv"

	video_count "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/count"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/zeromicro/go-zero/core/threading"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPlayHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPlayHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPlayHistoryLogic {
	return &AddPlayHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddPlayHistory 添加播放历史
func (l *AddPlayHistoryLogic) AddPlayHistory(req *types.PlayHistoryReq) (resp *types.PlayHistoryResp, err error) {
	resp = &types.PlayHistoryResp{}
	doerId := jwt.GetUidFromCtx(l.ctx)
	videoId, err := strconv.ParseInt(req.VideoId, 10, 64)
	if err != nil {
		return nil, err
	}
	// 用户增加播放历史
	err = l.svcCtx.VideoHistoryModel.AddCachePlayHistory(l.ctx, doerId, videoId) // TODO:处理用户重复播放
	if err != nil {
		return nil, err
	}
	threading.GoSafe(func() { // 视频增加播放量
		_, err := video_count.NewVideoCountModel(l.svcCtx.BizRedis).IncrVideoPlayCount(l.ctx, videoId)
		if err != nil {
			logx.Errorf("[AddPlayHistory] VideoRpc.IncrVideoPlayCount error: %v", err)
		}
		// TODO:消息队列将播放量同步到数据库
	})
	return
}
