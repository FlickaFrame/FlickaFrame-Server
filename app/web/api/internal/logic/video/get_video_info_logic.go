package video

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/pkg/jwt"

	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVideoInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoInfoLogic {
	return &GetVideoInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVideoInfoLogic) GetVideoInfo(req *types.GetVideoInfoReq) (resp *types.GetVideoInfoResp, err error) {
	video, err := l.svcCtx.VideoModel.FindOne(l.ctx, req.VideoId)
	if err != nil {
		logx.Info(err)
		return nil, err
	}
	videoBasicInfo, err := NewConvert(l.ctx, l.svcCtx).BuildVideoBasicInfo(l.ctx, video)
	resp = &types.GetVideoInfoResp{
		Video: videoBasicInfo,
	}
	// 互动信息(与当前登录用户的关注关系)
	doerId := jwt.GetUidFromCtx(l.ctx)
	resp.Video.VideoUserInfo.IsFollow = l.svcCtx.UserModel.IsFollowing(l.ctx, doerId, video.AuthorID)
	return
}
