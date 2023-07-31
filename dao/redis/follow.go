package redis

import (
	"database/sql"

	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/model"
	"go.uber.org/zap"
)

func Follow(uid, toUid string, attitude bool) (err error) {
	// TODO 添加关注时，如果对方已经关注了自己，则为互相关注
	pipeline := rdb.TxPipeline()
	if attitude {
		rdb.SAdd(ctx, getRedisKey(FollowerSetPrefix+toUid), uid)
		rdb.SAdd(ctx, getRedisKey(FollowingSetPrefix+uid), toUid)
	} else {
		rdb.SRem(ctx, getRedisKey(FollowerSetPrefix+toUid), uid)
		rdb.SRem(ctx, getRedisKey(FollowingSetPrefix+uid), toUid)
	}
	_, err = pipeline.Exec(ctx)

	return
}

func GetFollowers(uid string) (users []*model.UserDetail, err error) {
	// redis中获取id列表
	key := getRedisKey(FollowerSetPrefix) + uid
	ids, err := rdb.SMembers(ctx, key).Result()
	if err != nil {
		zap.L().Error("get ids from redis err", zap.Error(err))
		return nil, err
	}
	if len(ids) == 0 {
		err = sql.ErrNoRows
		zap.L().Error("no data")
		return nil, err
	}
	// 从数据库查询
	users, err = mysql.GetUsersByIds(ids)
	//PostIdStrings, err = rdb.ZRevRange(ctx, key, (form.Page-1)*form.Size, form.Page*form.Size-1).Result()
	return
}

func GetFollowings(uid string) (users []*model.UserDetail, err error) {
	// redis中获取id列表
	key := getRedisKey(FollowingSetPrefix) + uid
	ids, err := rdb.SMembers(ctx, key).Result()
	if err != nil {
		zap.L().Error("get ids from redis err", zap.Error(err))
		return nil, err
	}
	if len(ids) == 0 {
		err = sql.ErrNoRows
		zap.L().Error("no data")
		return nil, err
	}
	// 从数据库查询
	users, err = mysql.GetUsersByIds(ids)
	//PostIdStrings, err = rdb.ZRevRange(ctx, key, (form.Page-1)*form.Size, form.Page*form.Size-1).Result()
	return
}

func GetFriends(uid string) (users []*model.UserDetail, err error) {
	// redis中获取id列表
	//key := getRedisKey(FollowerSetPrefix) + uid
	//ids, err := rdb.SMembers(ctx, key).Result()
	//rdb.SInter()
	//if err != nil {
	//	zap.L().Error("get ids from redis err", zap.Error(err))
	//	return nil, err
	//}
	//// 从数据库查询
	////
	//users, err = mysql.GetUsersByIds(ids)
	//PostIdStrings, err = rdb.ZRevRange(ctx, key, (form.Page-1)*form.Size, form.Page*form.Size-1).Result()
	return
}
