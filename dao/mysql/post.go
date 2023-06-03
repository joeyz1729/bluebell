package mysql

import (
	"zouyi/bluebell/model"
)

func CreatePost(p *model.Post) (err error) {
	sqlStr := `insert into post(post_id, author_id, community_id, title, content) values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.Id, p.AuthorId, p.CommunityId, p.Title, p.Content)
	if err != nil {
		return ErrInsertPost
	}

	return
}

func GetPostList(page, size uint64) (posts []*model.Post, err error) {

	posts = make([]*model.Post, 0, size)
	sqlStr := `select post_id, author_id, community_id, title, content, create_time from post order by update_time limit ? offset ? `
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
