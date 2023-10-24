package config

import (
	"github.com/FlickaFrame/FlickaFrame-Server/pkg/orm"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	Oss struct { // 七牛云OSS配置
		Endpoint         string
		AccessKeyId      string
		AccessKeySecret  string
		BucketName       string
		ConnectTimeout   int64 `json:",optional"`
		ReadWriteTimeout int64 `json:",optional"`
	}
	Mysql    orm.Config
	BizRedis redis.RedisConf
}
