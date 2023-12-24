package svc

import (
	follow_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/follow/rpc/follow"
	user_rpc "github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/config"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/comment"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/favorite"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/notice"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video"
	video_history "github.com/FlickaFrame/FlickaFrame-Server/app/web/api/internal/model/video/history"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/go-playground/validator/v10"
	"github.com/meilisearch/meilisearch-go"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	Validate          *validator.Validate // 入参校验器
	DB                *orm.DB             // 数据库连接
	BizRedis          *redis.Redis        // 业务redis连接
	VideoModel        *video.VideoModel
	VideoHistoryModel *video_history.VideoHistoryModel
	UserModel         *user.UserModel
	FavoriteModel     *favorite.Model
	CommentModel      *comment.Model
	NoticeModel       *notice.NoticeModel
	Indexer           *meilisearch.Client
	KqPusherClient    *kq.Pusher
	UserRpc           user_rpc.User
	FollowRpc         follow_rpc.Follow
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.Mysql.DSN,
		MaxOpenConns: c.Mysql.MaxOpenConns,
		MaxIdleConns: c.Mysql.MaxIdleConns,
		MaxLifetime:  c.Mysql.MaxLifetime,
	})

	err := model.Migrate(db.DB)
	if err != nil {
		panic(err)
	}

	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	indexer := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   c.MeiliSearch.Host,
		APIKey: c.MeiliSearch.APIKey,
		//Timeout: time.Millisecond*c.MeiliSearch.Timeout
	})
	return &ServiceContext{
		Config:   c,
		Validate: validator.New(),

		UserRpc:           user_rpc.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		FollowRpc:         follow_rpc.NewFollow(zrpc.MustNewClient(c.FollowRpcConf)),
		DB:                db,
		BizRedis:          rds,
		VideoModel:        video.NewVideoModel(db),
		VideoHistoryModel: video_history.NewVideoHistoryModel(rds),
		UserModel:         user.NewUserModel(db, rds),
		FavoriteModel:     favorite.NewFavoriteModel(db),
		CommentModel:      comment.NewCommentModel(db),
		NoticeModel:       notice.NewNoticeModel(db),
		Indexer:           indexer,
		KqPusherClient:    kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
