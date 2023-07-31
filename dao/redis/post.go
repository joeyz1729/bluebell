package redis

import (
	"context"
	"strconv"
	"time"

	"github.com/YiZou89/bluebell/model"

	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

// CreatePost 将创建的帖子信息添加到redis中
func CreatePost(pid, cid uint64) (err error) {
	pipeline := rdb.TxPipeline()
	ctx := context.Background()

	// 按照创建时间排序 bluebell:post:time
	_, err = pipeline.ZAdd(ctx, getRedisKey(PostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: int64(pid),
	}).Result()

	// 按照投票数排序 bluebell:post:score
	_, err = pipeline.ZAdd(ctx, getRedisKey(PostScoreZSet), redis.Z{
		Score:  0,
		Member: int64(pid),
	}).Result()

	// 同社区内按照投票数排序 bluebell:community:cid
	cKey := getRedisKey(CommunityPostSetPrefix + strconv.Itoa(int(cid)))
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

// GetPostIdsInOrder 按照指定的排列顺序获取帖子分页信息
func GetPostIdsInOrder(order string, page, size int64) (PostIdStrings []string, err error) {
	key := getRedisKey(PostTimeZSet)
	if order == model.OrderByScore {
		key = getRedisKey(PostScoreZSet)
	}
	ctx := context.Background()
	PostIdStrings, err = rdb.ZRevRange(ctx, key, (page-1)*size, size-1).Result()

	return
}

// GetPostVoteData 按照id列表从redis中获取投票数
func GetPostVoteData(ids []string) (data []int64, err error) {
	ctx := context.Background()
	pipeline := rdb.TxPipeline()
	for _, pid := range ids {
		// bluebell:post:voted:pid
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

	cKey := getRedisKey(CommunityPostSetPrefix) + strconv.Itoa(int(form.CommunityId))
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

// getIdsFormKey 获取指定分页内的记录
func getIdsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := page*size - 1
	return rdb.ZRevRange(context.Background(), key, start, end).Result()
}
