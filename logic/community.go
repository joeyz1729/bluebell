package logic

import (
	"errors"

	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/model"
	"go.uber.org/zap"
)

func CommunityJoin(uid uint64, jf *model.JoinForm) (err error) {
	// TODO， 添加到redis， 添加到mysql， 是否使用rabbitmq
	zap.L().Debug("[logic] community join")
	err = mysql.CheckIfJoin(jf.CommunityID, uid)
	if err == nil {
		// 有记录
		err = mysql.CommunityChangeJoin(jf.CommunityID, uid, jf.ActionType)
	} else if errors.Is(err, mysql.ErrorNotJoin) {
		// 没有记录
		err = mysql.CommunityJoin(jf.CommunityID, uid, jf.ActionType)
	}
	return
}

func GetCommunityList() (list []*model.Community, err error) {
	list, err = mysql.GetAllCommunityList()
	// 业务错误处理
	return
}

func GetCommunityJoinList(uid uint64) (list []*model.Community, err error) {
	list, err = mysql.GetCommunityJoinList(uid)
	// 业务错误处理
	return
}
