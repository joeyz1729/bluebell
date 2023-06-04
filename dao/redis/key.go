package redis

const (
	Prefix              = "bluebell:"
	PostTimeZSet        = "post:time"
	PostScoreZSet       = "post:score"
	PostVotedZSetPrefix = "post:voted:"
	CommunitySetPrefix  = "community:"
)

// community:cid set(pid)
// post:voted:pid uid ifVoted

func getRedisKey(key string) string {
	return Prefix + key
}
