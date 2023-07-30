package redis

import (
	"context"
	"fmt"

	"github.com/YiZou89/bluebell/setting"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var rdb *redis.Client

func Init(conf *setting.RedisConfig) (err error) {
	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize,
	})

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("connect to redis failed, err: %s\n", zap.Error(err))
		return
	}
	zap.L().Info("[redis] init success")
	return
}

func Close() {
	_ = rdb.Close()
}
