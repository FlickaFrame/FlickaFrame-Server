package config

import (
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	UserRpcConf   zrpc.RpcClientConf
	FollowRpcConf zrpc.RpcClientConf
	MeiliSearch   struct {
		Host    string
		APIKey  string
		Timeout int64
	}
	Mysql        orm.Config
	BizRedis     redis.RedisConf
	KqPusherConf struct {
		Brokers []string
		Topic   string
	}
	KqConsumerConf kq.KqConf
}
