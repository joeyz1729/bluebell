package redis

const (
	Prefix = "bluebell:"

	// 用于帖子排序功能
	PostTimeZSet  = "post:time"  // bluebell:post:time, Sorted Set, key为pid, value为计时
	PostScoreZSet = "post:score" // bluebell:post:score, Sorted Set, key为pid, value为投票得分

	PostVotedZSetPrefix       = "post:voted:"     // bluebell:post:voted:pid, Sorted Set, key为uid, value为投票内容(-1, 0, 1)
	CommunityPostSetPrefix    = "community:post:" // community:cid, Sorted Set, key为pid, value默认取0
	AuthorPostSetPrefix       = "work:post:"
	AuthorPostVotedZSetPrefix = "work:post:vote:" // 统计发帖数，总获赞数

	CommunityMemberPrefix = "community:member:"
	MemberCommunityPrefix = "member:community:"

	FollowerSetPrefix  = "follower:"
	FollowingSetPrefix = "following:"
)

// join 参加
// 用于判断用户是否加入社区
// set community:member:cid (key uid)

// post 帖子内容
// 统计发帖数
// set work:uid (key pid)
// 统计个人总获赞数
// zset work:post:uid (key pid: vote)
// score和就是总获赞，每个score就是pid的获赞
// 判断是否点赞, 因为有踩功能，所以使用zset
// zset post:voted:pid (key uid actionType)
//

// zset post:time (key pid, timestamp)
// zset post:vote (key pid, voteCnt)
// zset

// follow 关注
// 用于计算关注数，和被关注数
// 判断是否已关注
// set follower:userid (key follower id)
// set following: userid (key following id)
//

func getRedisKey(key string) string {
	return Prefix + key
}
