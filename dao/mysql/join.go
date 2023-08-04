package mysql

func CommunityJoin(cid, uid string, attitude bool) (err error) {
	var action int
	if !attitude {
		action = 1
	}
	sqlStr := `insert into member (community_id, user_id, cancel) VALUES (?, ?, ?)`
	_, err = db.Exec(sqlStr, cid, uid, action)
	return

}

func CommunityChangeJoin(cid, uid string, attitude bool) (err error) {
	var action int
	if !attitude {
		action = 1
	}
	sqlStr := `update member set cancel = ? where community_id = ? and user_id = ?`
	_, err = db.Exec(sqlStr, action, cid, uid)
	return
}

func CheckIfJoin(cid, uid string) (err error) {
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
