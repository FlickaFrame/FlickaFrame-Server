// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	feed "github.com/FlickaFrame/FlickaFrame-Server/internal/handler/feed"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/feed",
				Handler: feed.FeedHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}