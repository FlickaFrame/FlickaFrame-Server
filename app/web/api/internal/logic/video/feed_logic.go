package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/util"
	"github.com/zeromicro/go-zero/core/logx"
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
	videoConvert := NewConvert(l.ctx, l.svcCtx)
	return videoConvert.Feed(req, func(*svc.ServiceContext) ([]*video_model.Video, error) {
		return l.svcCtx.VideoModel.List(l.ctx, video_model.ListOption{
			AuthorID:   util.MustString2Int64(req.AuthorID),
			LatestTime: time.UnixMilli(req.Cursor),
			Limit:      req.Limit,
			QueryAll:   false,
			CategoryID: util.MustString2Int64(req.CategoryID),
		})
	})
}
