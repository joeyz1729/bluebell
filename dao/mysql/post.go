package mysql

import (
	"strings"

	"github.com/YiZou89/bluebell/model"

	"github.com/jmoiron/sqlx"
)

func CreatePost(p *model.Post) (err error) {
	sqlStr := `insert into post(post_id, author_id, community_id, title, content) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.Id, p.AuthorId, p.CommunityId, p.Title, p.Content)
	if err != nil {
		return ErrInsertPost
	}

	return
}

func GetPostList(page, size int64) (posts []*model.Post, err error) {

	posts = make([]*model.Post, 0, size)
	sqlStr := `select post_id, author_id, community_id, title, content, create_time from post order by update_time desc limit ? offset ? `
	err = db.Select(&posts, sqlStr, size, (page-1)*size)
	return
}

func GetPostById(pid uint64) (pd *model.Post, err error) {
	pd = new(model.Post)
	pd.Id = pid
	sqlStr := `select author_id, community_id, title, content, create_time from post where post_id = ?`
	err = db.Get(pd, sqlStr, pid)

	return
}

func GetPostByIds(ids []string) (posts []*model.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post	where post_id in (?) order by FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&posts, query, args...)
	return
}

func GetCommunityPostList(cid uint64, page, size int64) (posts []*model.Post, err error) {
	posts = make([]*model.Post, 0, size)
	sqlStr := `select post_id, author_id, title, content, create_time from post where community_id = ? order by update_time desc limit ? offset ?`
	err = db.Select(&posts, sqlStr, cid, size, (page-1)*size)
	return
}
