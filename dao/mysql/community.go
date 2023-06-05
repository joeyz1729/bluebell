package mysql

import (
	"database/sql"
	"zouyi/bluebell/model"

	"go.uber.org/zap"
)

var ()

func GetCommunityList() (communityList []*model.Community, err error) {
	sqlStr := `select community_id, community_name from community`
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		zap.L().Warn("no community in db")
		err = nil
		return
	}
	return
}

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
