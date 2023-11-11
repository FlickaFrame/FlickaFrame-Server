// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/handler/comment"
	favorite "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/handler/favorite"
	notice "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/handler/notice"
	oss "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/handler/oss"
	user "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/handler/user"
	video "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/handler/video"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/detail",
				Handler: user.CurrentUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/info",
				Handler: user.UpdateInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/updatepwd",
				Handler: user.UpdatePasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/detail/:userId",
				Handler: user.GetUserDetailInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/ranking",
				Handler: user.RankingHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/video/create",
				Handler: video.CreateVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/video/:videoId",
				Handler: video.DeleteVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/video/following",
				Handler: video.FollowingHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/video/detail/:videoId",
				Handler: video.GetVideoInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/video/category",
				Handler: video.CategoryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/video/feed",
				Handler: video.FeedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/video/search",
				Handler: video.SearchHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/video/play-history",
				Handler: video.AddPlayHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/video/play-history",
				Handler: video.GetPlayHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/video/play-history/:videoId",
				Handler: video.DeletePlayHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/video/play-history",
				Handler: video.ClearPlayHistoryHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/video/:video_id/comments/:comment_id",
				Handler: comment.GetVideoCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/video/:video_id/comments",
				Handler: comment.ListVideoCommentsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comment/video",
				Handler: comment.CreateVideoCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/parent",
				Handler: comment.CreateChildCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/comment/child",
				Handler: comment.CreateReplyCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/comment/:comment_id",
				Handler: comment.DeleteVideoCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/:comment_id",
				Handler: comment.EditVideoCommentHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/favorite/video/:targetId",
				Handler: favorite.FavoriteVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/favorite/video/:targetId",
				Handler: favorite.UnFavoriteCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favorite/video/:targetId",
				Handler: favorite.CheckVideoFavoriteHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/favorite/comment/:targetId",
				Handler: favorite.FavoriteCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/favorite/comment/:targetId",
				Handler: favorite.UnFavoriteVideoHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favorite/comment/:targetId",
				Handler: favorite.CheckCommentFavoriteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/oss/endpoint",
				Handler: oss.EndpointHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/oss/uptoken",
				Handler: oss.CreateUpTokenHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/notice/type",
				Handler: notice.GetNoticeHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/user/follow_action/:user_id",
				Handler: user.FollowHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/follow_action/:user_id",
				Handler: user.UnfollowHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/me/followers",
				Handler: user.ListMyFollowersHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/me/following",
				Handler: user.ListMyFollowingHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/:user_id/followers",
				Handler: user.ListFollowersHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user/:user_id/following",
				Handler: user.ListFollowingHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
