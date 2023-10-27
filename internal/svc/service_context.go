package svc

import (
	"github.com/FlickaFrame/FlickaFrame-Server/internal/config"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/favorite"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/user"
	"github.com/FlickaFrame/FlickaFrame-Server/internal/model/video"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/go-playground/validator/v10"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config        config.Config
	Validate      *validator.Validate    // 入参校验器
	UploadManager *storage.UploadManager // 七牛云上传管理器
	DB            *orm.DB                // 数据库连接
	BizRedis      *redis.Redis           // 业务redis连接
	VideoModel    *video.VideoModel
	UserModel     *user.UserModel
	FollowModel   *user.FollowModel
	FavoriteModel *favorite.FavoriteModel
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

	return &ServiceContext{
		Config:   c,
		Validate: validator.New(),
		UploadManager: storage.NewUploadManager(&storage.UploadConfig{
			UseHTTPS:      true,
			UseCdnDomains: false,
		}),
		DB:            db,
		BizRedis:      rds,
		VideoModel:    video.NewVideoModel(db.DB),
		UserModel:     user.NewUserModel(db.DB),
		FollowModel:   user.NewFollowModel(db.DB),
		FavoriteModel: favorite.NewFavoriteModel(db.DB),
	}
}
