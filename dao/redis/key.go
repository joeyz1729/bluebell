package redis

const (
	Prefix              = "bluebell:"
	PostTimeZSet        = "post:time"   // bluebell:post:time, Sorted Set, key为pid, value为计时
	PostScoreZSet       = "post:score"  // bluebell:post:score, Sorted Set, key为pid, value为投票得分
	PostVotedZSetPrefix = "post:voted:" // bluebell:post:voted:pid, Sorted Set, key为uid, value为投票内容(-1, 0, 1)
	CommunitySetPrefix  = "community:"  // community:cid, Sorted Set, key为pid, value默认取0
)

func getRedisKey(key string) string {
	return Prefix + key
}
