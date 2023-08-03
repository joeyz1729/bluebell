package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/YiZou89/bluebell/setting"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	oneWeek = 7 * 24 * 3600 * time.Second
)

var (
	ErrVoteTimeExpire         = errors.New("vote time expired")
	ErrVoteRepeated           = errors.New("vote repeated")
	voteScore         float64 = 432
	ctx                       = context.Background()
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
	return nil
}

func Close() {
	_ = rdb.Close()
}
