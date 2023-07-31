package mysql

func CommunityJoin(cid, uid uint64, action int32) (err error) {
	if action == 2 {
		action = 0
	}
	sqlStr := `insert into member (community_id, user_id, cancel) VALUES (?, ?, ?)`
	_, err = db.Exec(sqlStr, cid, uid, action)
	return

}

func CommunityChangeJoin(cid, uid uint64, action int32) (err error) {
	if action == 2 {
		action = 0
	}
	sqlStr := `update member set cancel = ? where community_id = ? and user_id = ?`
	_, err = db.Exec(sqlStr, action, cid, uid)
	return
}

func CheckIfJoin(cid, uid uint64) (err error) {
	sqlStr := `select count(id) from member where user_id = ? and community_id = ?`
	var count int
	err = db.Get(&count, sqlStr, uid, cid)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrorNotJoin
	}
	return
}
