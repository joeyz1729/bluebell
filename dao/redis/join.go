package redis

import (
	"context"
)

//
//CommunityMemberPrefix = "community:member:"
//MemberCommunitySetPrefix = "member:community:"

func JoinCommunity(uid, cid string, action bool) (err error) {
	// 有两个缓存集合，使用事务
	pipeline := rdb.TxPipeline()
	if action {
		rdb.SAdd(context.Background(), getRedisKey(CommunityMemberSetPrefix)+cid, uid)
		rdb.SAdd(context.Background(), getRedisKey(MemberCommunitySetPrefix)+uid, cid)
	} else {
		rdb.SRem(context.Background(), getRedisKey(CommunityMemberSetPrefix)+cid, uid)
		rdb.SRem(context.Background(), getRedisKey(MemberCommunitySetPrefix)+uid, cid)
	}
	_, err = pipeline.Exec(context.Background())
	return
}

func IsMember(uid, cid string) (ok bool, err error) {
	ok, err = rdb.SIsMember(context.Background(), getRedisKey(CommunityMemberSetPrefix+cid), uid).Result()
	return
}

func DelJoinCommunity(uid, cid string) (err error) {
	return
}

func GetCommunityJoinList(uid string) (ids []string, err error) {
	return rdb.SMembers(context.Background(), getRedisKey(MemberCommunitySetPrefix+uid)).Result()

}
