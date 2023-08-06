package logic

import (
	"errors"
	"strconv"

	"github.com/YiZou89/bluebell/dao/redis"

	"go.uber.org/zap"

	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/model"
)

func CommunityJoin(uid, cid string, attitude bool) (err error) {
	// 更新数据库
	zap.L().Debug("[logic] community join")
	err = mysql.CheckIfJoin(cid, uid)
	if err == nil {
		// 有记录
		err = mysql.CommunityChangeJoin(cid, uid, attitude)
	} else if errors.Is(err, mysql.ErrorNotJoin) {
		// 没有记录
		err = mysql.CommunityJoin(cid, uid, attitude)
	}
	// 删除redis缓存
	return redis.DelJoinCommunity(uid, cid)

	// 检查redis是否有
	// 添加+有 不用管； 添加+无 要加； 删除+有 要删除； 删除+无 不用管
	// TODO，cache-aside写时用不用检查？
	//exists, err := redis.IsMember(uid, cid)
	//fmt.Println(exists)
	//if err != nil {
	//	zap.L().Error("check redis is member err:", zap.Error(err))
	//	return err
	//}
	//// 如果有，更新并返回，publish到消息队列
	//if exists {
	//	if err = redis.JoinCommunity(uid, cid, attitude); err != nil {
	//		// 添加缓存失败，删除并消息队列
	//		//redis.JoinCommunity(uid, cid, false)
	//		if err = redis.DelJoinCommunity(uid, cid); err != nil {
	//			return err
	//		}
	//	}
	//	// TODO 消息队列
	//	return
	//}

	// 如果没有，检查数据库

}

func GetCommunityList() (list []*model.Community, err error) {
	list, err = mysql.GetAllCommunityList()
	// 业务错误处理
	return
}

func GetCommunityJoinList(uid uint64) (list []*model.Community, err error) {
	// 问题：
	// 从 redis 获取community id， 再mysql获取，
	ids, err := redis.GetCommunityJoinList(strconv.Itoa(int(uid)))
	if err != nil {
		return
	}
	// 没有缓存
	if len(ids) == 0 {
		list, err = mysql.GetCommunityJoinList(uid)
		for _, community := range list {
			cidStr, uidStr := strconv.Itoa(int(community.CommunityId)), strconv.Itoa(int(uid))
			// TODO redis事务添加？`
			mysql.CommunityJoin(cidStr, uidStr, true)
		}
		return
	}
	list, err = mysql.GetCommunityByIds(ids)
	return
}
