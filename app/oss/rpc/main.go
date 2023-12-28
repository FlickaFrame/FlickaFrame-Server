package main

import (
	"flag"
	"fmt"

	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/config"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/server"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/internal/svc"
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/pb/oss_service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/main.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		oss_service.RegisterOssServer(grpcServer, server.NewOssServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
