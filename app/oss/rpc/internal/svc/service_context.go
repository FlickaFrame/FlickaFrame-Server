package svc

import "github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
