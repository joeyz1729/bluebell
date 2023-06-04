package redis

import (
	"context"
	"strconv"
	"time"
	"zouyi/bluebell/model"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

func CreatePost(pid, cid uint64) (err error) {
	pipeline := rdb.TxPipeline()
	ctx := context.Background()
	// add post create time
	_, err = pipeline.ZAdd(ctx, getRedisKey(PostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: int64(pid),
	}).Result()

	// add post vote score
	_, err = pipeline.ZAdd(ctx, getRedisKey(PostScoreZSet), redis.Z{
		Score:  0,
		Member: int64(pid),
	}).Result()

	// add post community id
	cKey := getRedisKey(CommunitySetPrefix + strconv.Itoa(int(cid)))
	_, err = pipeline.ZAdd(ctx, cKey, redis.Z{
		Score:  0,
		Member: int64(pid),
	}).Result()

	_, err = pipeline.Exec(ctx)
	if err != nil {
		zap.L().Error("redis pipeline create post err", zap.Error(err))
		return
	}
	return
}

func GetPostIdsInOrder(form *model.PostsForm) (PostIdStrings []string, err error) {
	key := getRedisKey(PostTimeZSet)
	if form.Order == model.OrderByScore {
		key = getRedisKey(PostScoreZSet)
	}
	ctx := context.Background()
	PostIdStrings, err = rdb.ZRevRange(ctx, key, (form.Page-1)*form.Size, form.Page*form.Size-1).Result()

	return
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	ctx := context.Background()
	pipeline := rdb.TxPipeline()
	for _, pid := range ids {
		key := getRedisKey(PostVotedZSetPrefix + pid)
		pipeline.ZCount(ctx, key, "1", "1")
	}
	cmders, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}

func GetCommunityPostIdsInOrder(form *model.CommunityPostsForm) (PostIdStrings []string, err error) {
	orderKey := getRedisKey(PostTimeZSet)
	if form.Order == model.OrderByScore {
		orderKey = getRedisKey(PostScoreZSet)
	}

	cKey := getRedisKey(CommunitySetPrefix) + strconv.Itoa(int(form.CommunityId))
	key := orderKey + strconv.Itoa(int(form.CommunityId))
	ctx := context.Background()

	if rdb.Exists(ctx, key).Val() < 1 {
		pipeline := rdb.TxPipeline()
		pipeline.ZInterStore(ctx, key, &redis.ZStore{
			Keys:      []string{cKey, orderKey},
			Aggregate: "MAX",
		})
		pipeline.Expire(ctx, key, 60*time.Second)
		_, err = pipeline.Exec(ctx)
		if err != nil {
			return nil, err
		}
	}
	return getIdsFormKey(key, form.Page, form.Size)

}

func getIdsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := page*size - 1
	return rdb.ZRevRange(context.Background(), key, start, end).Result()
}
