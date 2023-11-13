package mqs

import (
	"context"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/config"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, NewActionSuccess(ctx, svcContext)),

		//.....
	}

}
