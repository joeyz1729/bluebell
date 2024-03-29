package mysql

func AddFollow(uid, toUid int) error {
	sqlStr := `insert into follow(user_id, follower_id, cancel) values(?, ?,?)`
	_, err := db.Exec(sqlStr, uid, toUid, 0)
	return err
}

func IsFollowed(uid, userId uint64) (ok bool, err error) {

	sqlStr := `select count(*) from follow where user_id = ? and follower_id = ?`
	var n int64
	err = db.Select(&n, sqlStr, uid, userId)
	if err != nil {
		return
	}
	return n != 0, nil
}

func GetFollowerCount(uid uint64) (n int64, err error) {
	sqlStr := `select count(*) from follow where user_id = ?`
	err = db.Select(&n, sqlStr, uid)
	return
}

func GetFollowingCount(uid uint64) (n int64, err error) {
	sqlStr := `select count(*) from follow where follower_id = ?`
	err = db.Select(&n, sqlStr, uid)
	return
}
