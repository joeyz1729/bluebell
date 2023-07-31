package redis

import "github.com/YiZou89/bluebell/model"

func GetUserInfo(uid, userId string) (info *model.UserDetail, err error) {
	// 关注总数， 粉丝总数， 发帖总数， 点赞总数，获赞总数
	info = new(model.UserDetail)

	// uid 是否关注userId
	iff, err := rdb.SIsMember(ctx, getRedisKey(FollowingSetPrefix)+uid, userId).Result()
	if err != nil {
		return
	}
	info.IsFollow = iff

	// 关注数
	fgc, err := rdb.SCard(ctx, getRedisKey(FollowingSetPrefix)+uid).Result()
	if err != nil {
		return
	}
	info.FollowCount = fgc

	// 粉丝数
	frc, err := rdb.SCard(ctx, getRedisKey(FollowerSetPrefix)+uid).Result()
	if err != nil {
		return
	}
	info.FollowerCount = frc

	// 加入的社区总数
	jc, err := rdb.SCard(ctx, getRedisKey(MemberCommunityPrefix)+uid).Result()
	if err != nil {
		return
	}
	info.JoinedCount = jc

	// 发布的帖子数
	wc, err := rdb.ZCard(ctx, getRedisKey(AuthorPostSetPrefix+uid)).Result()
	if err != nil {
		return
	}
	info.WorkCount = wc

	// 总获赞数，获取每个帖子的赞数求和
	var tf int64
	membersWithScore, err := rdb.ZRangeWithScores(ctx, getRedisKey(AuthorPostVotedZSetPrefix)+uid, 0, -1).Result()
	if err != nil {
		return
	}
	for _, member := range membersWithScore {
		tf += int64(member.Score)
	}
	//
	info.TotalFavorited = tf

	return info, nil
}
