package redis

import (
	"context"
	"errors"
	"math"
	"time"

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
)

func VoteForPost(userId, postId string, attitude float64) (err error) {
	var ctx = context.Background()
	// 1. check if the post has expired
	postTime := rdb.ZScore(ctx, getRedisKey(PostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > float64(oneWeek) {
		zap.L().Error("the post has been expired, can not be voted", zap.Error(err))
		return ErrVoteTimeExpire
	}

	// 2. update score in redis

	// check if the user vote on post
	voteValue := rdb.ZScore(ctx, getRedisKey(PostVotedZSetPrefix+postId), userId).Val()
	if attitude == voteValue {
		return ErrVoteRepeated
	}

	var sign float64
	sign = -1
	if attitude > voteValue {
		sign = 1
	}
	diff := math.Abs(voteValue - attitude)

	// begin transaction
	pipeline := rdb.TxPipeline()
	// update post vote score
	_, err = pipeline.ZIncrBy(ctx, getRedisKey(PostScoreZSet), sign*diff*voteScore, postId).Result()

	// update post-user vote
	if attitude == 0 {
		_, err = pipeline.ZRem(ctx, getRedisKey(PostVotedZSetPrefix)+postId, userId).Result()

	} else {
		_, err = pipeline.ZAdd(ctx, getRedisKey(PostVotedZSetPrefix+postId), redis.Z{
			Score:  attitude,
			Member: userId,
		}).Result()

	}
	_, err = pipeline.Exec(ctx)
	// end transaction

	return
}
