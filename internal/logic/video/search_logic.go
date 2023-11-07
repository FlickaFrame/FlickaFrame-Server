package video

import (
	"context"
	video_model "github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/pkg/jwt"
	"github.com/jinzhu/copier"
	"github.com/meilisearch/meilisearch-go"
	"strconv"

	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.VideoSearchReq) (resp *types.VideoSearchResp, err error) {
	resp = &types.VideoSearchResp{}
	doerId := jwt.GetUidFromCtx(l.ctx)
	searchRes, err := l.svcCtx.Indexer.Index("video").Search(req.Keyword, &meilisearch.SearchRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	// 搜索结果
	err = copier.Copy(resp, searchRes)
	if err != nil {
		logx.Info(err)
		return nil, err
	}
	// 搜索结果转换
	videos := make([]*video_model.Video, 0, len(searchRes.Hits))
	for _, hit := range searchRes.Hits {
		id := hit.(map[string]interface{})["id"].(float64)
		video, err := l.svcCtx.VideoModel.FindOne(l.ctx, int64(id))
		if err != nil {
			logx.Info(err)
			return nil, err
		}
		videos = append(videos, video)
	}
	list, err := NewConvert(l.ctx, l.svcCtx).BuildVideoBasicInfoList(l.ctx, videos)
	if err != nil {
		return nil, err
	}
	resp.Videos = make([]*types.FeedVideoItem, len(list))
	err = copier.Copy(&resp.Videos, &list)
	// 判断关注状态
	for i := range list {
		authorId, _ := strconv.ParseInt(list[i].VideoUserInfo.ID, 10, 64)
		resp.Videos[i].VideoUserInfo.IsFollow = l.svcCtx.UserModel.IsFollowing(l.ctx, doerId, authorId)
	}
	return
}
