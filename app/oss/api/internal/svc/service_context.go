package svc

import (
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/api/internal/config"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/oss"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	OssRpc oss.Oss
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		OssRpc: oss.NewOss(zrpc.MustNewClient(c.OssRpcConf)),
	}
}
