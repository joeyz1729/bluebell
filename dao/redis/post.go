package redis

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

func CreatePost(postId uint64) (err error) {
	pipeline := rdb.TxPipeline()
	ctx := context.Background()
	// add post create time
	_, err = pipeline.ZAdd(ctx, getRedisKey(PostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: int64(postId),
	}).Result()

	// add post vote score
	_, err = pipeline.ZAdd(ctx, getRedisKey(PostScoreZSet), redis.Z{
		Score:  0,
		Member: int64(postId),
	}).Result()
	_, err = pipeline.Exec(ctx)
	if err != nil {
		zap.L().Error("redis pipeline create post err", zap.Error(err))
		return
	}
	return
}
