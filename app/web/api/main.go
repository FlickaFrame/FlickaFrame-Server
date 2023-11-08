package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/config"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/handler"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/middleware"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/mqs"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"

	"github.com/FlickaFrame/FlickaFrame-Server/pkg/open_api"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/xcode"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "app/web/api/etc/main.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf,
		rest.WithUnauthorizedCallback(xcode.UnAuthorizedCallback),
		rest.WithCors(),
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 对外开放Swagger
	open_api.RegisterSwagger(server)
	server.Use(middleware.NewCurrentUserMiddleware(c.JwtAuth.AccessSecret).Handle)

	// 自定义错误处理方法
	httpx.SetErrorHandler(xcode.ErrHandler)
	// 自定义成功处理方法
	httpx.SetOkHandler(xcode.OkHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// defer server.Start()

	// 启动消费者
	ctxbg := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	serviceGroup.Add(server)
	for _, mq := range mqs.Consumers(c, ctxbg, ctx) {
		serviceGroup.Add(mq)
	}
	serviceGroup.Start()
}
