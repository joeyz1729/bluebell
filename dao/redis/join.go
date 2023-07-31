package redis

import (
	"context"
	"time"

	"go.uber.org/zap"
)

func JoinCommunity(userId, postId string, action bool) (err error) {
	// action: 加入或者退出
	var ctx = context.Background()
	// 检查是否过期
	postTime := rdb.ZScore(ctx, getRedisKey(PostTimeZSet), postId).Val()
	if float64(time.Now().Unix())-postTime > float64(oneWeek) {
		zap.L().Error("the post has been expired, can not be voted", zap.Error(err))
		return ErrVoteTimeExpire
	}

	//// 获取uid为pid的投票值, -1, 0, 1
	//voteValue := rdb.ZScore(ctx, getRedisKey(PostVotedZSetPrefix+postId), userId).Val()
	//if attitude == voteValue {
	//	return ErrVoteRepeated
	//}
	//
	//var sign float64
	//sign = -1
	//if attitude > voteValue {
	//	sign = 1
	//}
	//diff := math.Abs(voteValue - attitude)
	//
	//// 修改
	//pipeline := rdb.TxPipeline()
	//_, err = pipeline.ZIncrBy(ctx, getRedisKey(PostScoreZSet), sign*diff*voteScore, postId).Result()
	//if attitude == 0 {
	//	_, err = pipeline.ZRem(ctx, getRedisKey(PostVotedZSetPrefix)+postId, userId).Result()
	//} else {
	//	_, err = pipeline.ZAdd(ctx, getRedisKey(PostVotedZSetPrefix+postId), redis.Z{
	//		Score:  attitude,
	//		Member: userId,
	//	}).Result()
	//}
	//_, err = pipeline.Exec(ctx)

	return
}
