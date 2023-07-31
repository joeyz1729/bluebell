package redis

import "github.com/YiZou89/bluebell/model"

func GetUserInfo(uid string) (info *model.UserDetail, err error) {
	// 关注总数， 粉丝总数， 发帖总数， 点赞总数，获赞总数

	// 关注数
	fgc, err := rdb.SCard(ctx, getRedisKey(FollowingSetPrefix)+uid).Result()
	if err != nil {
		return
	}

	// 粉丝数
	frc, err := rdb.SCard(ctx, getRedisKey(FollowerSetPrefix)+uid).Result()
	if err != nil {
		return
	}

	// 发布的帖子数
	wc, err := rdb.ZCard(ctx, getRedisKey(AuthorPostSetPrefix+uid)).Result()
	if err != nil {
		return
	}

	// 加入的社区总数
	jc, err := rdb.SCard(ctx, getRedisKey(MemberCommunityPrefix)+uid).Result()
	if err != nil {
		return
	}

	// 总获赞数，获取每个帖子的赞数求和
	var tf int64
	membersWithScore, err := rdb.ZRangeWithScores(ctx, getRedisKey(AuthorPostVotedZSetPrefix)+uid, 0, -1).Result()
	if err != nil {
		return
	}
	for _, member := range membersWithScore {
		tf += int64(member.Score)
	}
	info.FollowCount = fgc
	info.FollowerCount = frc
	info.WorkCount = wc
	info.TotalFavorited = tf
	info.JoinedCount = jc
	return info, nil
}
