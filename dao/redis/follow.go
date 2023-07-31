package redis

func Follow(uid, toUid string, attitude bool) (err error) {

	pipeline := rdb.TxPipeline()
	if attitude {
		rdb.SAdd(ctx, getRedisKey(FollowerSetPrefix+toUid), uid)
		rdb.SAdd(ctx, getRedisKey(FollowingSetPrefix+uid), toUid)
	} else {
		rdb.SRem(ctx, getRedisKey(FollowerSetPrefix+toUid), uid)
		rdb.SRem(ctx, getRedisKey(FollowingSetPrefix+uid), toUid)
	}
	_, err = pipeline.Exec(ctx)

	return
}
