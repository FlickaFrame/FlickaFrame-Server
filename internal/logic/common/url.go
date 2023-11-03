package common

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
)

type URLLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewURLLogic(ctx context.Context, svcCtx *svc.ServiceContext) *URLLogic {
	return &URLLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *URLLogic) GetAccessUrl(ctx context.Context, key string) string {
	return storage.MakePublicURL(c.svcCtx.Config.Oss.Endpoint, key)
}
