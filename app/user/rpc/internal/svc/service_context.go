package svc

import (
	"github.com/FlickaFrame/FlickaFrame-Server/app/oss/rpc/oss"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/config"
	"github.com/FlickaFrame/FlickaFrame-Server/app/user/rpc/internal/model"
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
	OssRpc    oss.Oss
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.Mysql.DSN,
		MaxOpenConns: c.Mysql.MaxOpenConns,
		MaxIdleConns: c.Mysql.MaxIdleConns,
		MaxLifetime:  c.Mysql.MaxLifetime,
	})
	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(db, rds),
		OssRpc:    oss.NewOss(zrpc.MustNewClient(c.OssRpcConf)),
	}
}
