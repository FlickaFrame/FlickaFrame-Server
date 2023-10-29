package video

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/meilisearch/meilisearch-go"

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

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchResp, err error) {
	searchRes, err := l.svcCtx.Indexer.Index("video").Search(req.Keyword, &meilisearch.SearchRequest{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	resp = &types.SearchResp{}
	err = copier.Copy(resp, searchRes)
	return
}
