package logic

import (
	"strconv"

	"github.com/YiZou89/bluebell/dao/redis"
	"github.com/YiZou89/bluebell/model"

	"go.uber.org/zap"
)

// follow关系
// uid, toUid, ifCancel存两个redis表， 一个mysql表
//

func Follow(uid, toUid uint64, attitude bool) (err error) {
	//TODO 拓展mysql和rabbitmq
	zap.L().Debug("logic.follow",
		zap.Uint64("userId", uid),
		zap.Uint64("postId", toUid),
		zap.Bool("attitude", attitude),
	)
	return redis.Follow(strconv.Itoa(int(uid)), strconv.Itoa(int(toUid)), attitude)

}

func GetFollowers(uid uint64) (users []*model.UserDetail, err error) {
	zap.L().Debug("logic.getFollowers",
		zap.Uint64("userId", uid),
	)
	users, err = redis.GetFollowers(strconv.Itoa(int(uid)))
	return
}

func GetFollowings(uid uint64) (users []*model.UserDetail, err error) {
	zap.L().Debug("logic.getFollowings",
		zap.Uint64("userId", uid),
	)
	users, err = redis.GetFollowings(strconv.Itoa(int(uid)))

	return
}

func GetFriends(uid uint64) (users []*model.UserDetail, err error) {
	zap.L().Debug("logic.getFriends",
		zap.Uint64("userId", uid),
	)
	users, err = redis.GetFriends(strconv.Itoa(int(uid)))
	return

}
