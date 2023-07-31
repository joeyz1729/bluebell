// follow关系
// uid, toUid, ifCancel存两个redis表， 一个mysql表
//

package logic

import (
	"strconv"

	"github.com/YiZou89/bluebell/dao/redis"
	"github.com/YiZou89/bluebell/model"

	"go.uber.org/zap"
)

func Follow(uid, toUid uint64, attitude bool) (err error) {
	zap.L().Debug("logic.post vote",
		zap.Uint64("userId", uid),
		zap.Uint64("postId", toUid),
		zap.Bool("attitude", attitude),
	)
	return redis.Follow(strconv.Itoa(int(uid)), strconv.Itoa(int(toUid)), attitude)

}

func GetFollowerList(uid uint64) (users []*model.UserDetail, err error) {

	return
}

func GetFollowingList(uid uint64) (users []*model.UserDetail, err error) {

	return
}

func GetFriendList(uid uint64) (users []*model.UserDetail, err error) {

	return
}
