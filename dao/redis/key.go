package redis

const (
	Prefix              = "bluebell:"
	PostTimeZSet        = "post:time"
	PostScoreZSet       = "post:score"
	PostVotedZSetPrefix = "post:voted:"
)

func getRedisKey(key string) string {
	return Prefix + key
}
