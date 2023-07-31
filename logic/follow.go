package logic

import "github.com/YiZou89/bluebell/model"

// follow关系
// uid, toUid, ifCancel存两个redis表， 一个mysql表
//

func Follow(uid, toUid uint64, action int64) (err error) {

	return
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
