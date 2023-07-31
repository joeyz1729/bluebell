package mysql

import (
	"database/sql"

	"github.com/YiZou89/bluebell/model"

	"go.uber.org/zap"
)

// GetAllCommunityList 获取所有社区信息
func GetAllCommunityList() (communityList []*model.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("no community in db")
		err = nil
		return
	}
	return
}

func GetCommunityJoinList(uid uint64) (list []*model.Community, err error) {
	//sqlStr := `select member.community_id, community_name from member, community where member.user_id = ? and member.community_id = community.community_id`
	sqlStr := `select community_id from member where user_id = ?`
	err = db.Select(&list, sqlStr, uid)
	if err == sql.ErrNoRows {
		zap.L().Warn("The user didn't join any community")
		err = nil
		return
	}
	return
}

// GetCommunityDetailById 获取指定社区信息
func GetCommunityDetailById(communityId uint64) (cd *model.CommunityDetail, err error) {
	cd = new(model.CommunityDetail)
	cd.CommunityId = communityId
	sqlStr := `select community_name, introduction, create_time from community where community_id = ?`
	err = db.Get(cd, sqlStr, communityId)
	if err == sql.ErrNoRows {
		err = ErrInvalidId
		return
	}
	if err != nil {
		err = ErrQueryFailed
		return
	}
	return
}

func GetCommunityNameById(communityId uint64) (cn string, err error) {

	sqlStr := `select community_name from community where community_id = ?`
	err = db.Get(cn, sqlStr)
	if err == sql.ErrNoRows {
		err = ErrInvalidId
		return
	}
	if err != nil {
		err = ErrQueryFailed
		return
	}
	return
}

func GetJoinCount(uid uint64) (n int64, err error) {
	sqlStr := `select count(*) from member where user_id = ?`
	err = db.Get(&n, sqlStr, uid)
	return
}
