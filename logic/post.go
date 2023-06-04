package logic

import (
	"zouyi/bluebell/dao/mysql"
	"zouyi/bluebell/dao/redis"
	"zouyi/bluebell/model"
	"zouyi/bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *model.Post) (err error) {
	// 1. param
	// post id,
	id, err := snowflake.GenID()
	if err != nil {
		return
	}
	p.Id = id
	// 2. save to redis
	err = redis.CreatePost(p.Id)
	if err != nil {
		zap.L().Error("post info save to redis err", zap.Error(err))
		return
	}

	// 2. save to database
	err = mysql.CreatePost(p)
	if err != nil {
		zap.L().Error("post info save to mysql err", zap.Error(err))
		return
	}
	return
}

func GetPostDetailById(pid uint64) (pd *model.PostDetail, err error) {
	pd = new(model.PostDetail)
	// 1. get post
	var post = new(model.Post)
	post, err = mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("GetPostById() failed", zap.Error(err))
		return
	}
	pd.Post = post
	// 2. get username
	user, err := mysql.GetUserById(pd.AuthorId)
	if err != nil {
		zap.L().Error("GetUserById() failed", zap.Error(err))
		return
	}
	pd.AuthorName = user.Username
	// 3. get community detail
	community, err := mysql.GetCommunityDetailById(pd.CommunityId)
	if err != nil {
		zap.L().Error("GetCommunityDetailById() failed", zap.Error(err))
		return
	}
	pd.CommunityDetail = community

	return
}

func GetPostList(page, size int64) (postDetailList []*model.PostDetail, err error) {
	postDetailList = make([]*model.PostDetail, 0, size)
	var posts []*model.Post
	posts, err = mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("GetPostList() err", zap.Error(err))
		return nil, err
	}
	for _, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post

		// get username by author_id
		var user = new(model.User)
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("GetUserById err", zap.Error(err))
			return
		}
		postDetail.AuthorName = user.Username

		// get communityDetail by community_id
		var communityDetail = new(model.CommunityDetail)
		communityDetail, err = mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("GetCommunityDetailById err", zap.Error(err))
			return
		}
		postDetail.CommunityDetail = communityDetail

		postDetailList = append(postDetailList, postDetail)
	}
	return
}

func GetPostListInOrder(form *model.PostListForm) (postDetailList []*model.PostDetail, err error) {
	postDetailList = make([]*model.PostDetail, 0, form.Size)
	var posts []*model.Post
	// get id in order in from redis
	ids, err := redis.GetPostIdsInOrder(form)
	if err != nil || len(ids) == 0 {
		zap.L().Error("get post ids in order from redis err, or len(ids) == 0", zap.Error(err))
		return
	}
	// get data from mysql
	posts, err = mysql.GetPostByIds(ids)
	if err != nil {
		zap.L().Error("GetPostListByIds from mysql err", zap.Error(err))
		return nil, err
	}
	// query post detail
	for _, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post

		// get username by author_id
		var user = new(model.User)
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("GetUserById err", zap.Error(err))
			return
		}
		postDetail.AuthorName = user.Username

		// get communityDetail by community_id
		var communityDetail = new(model.CommunityDetail)
		communityDetail, err = mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("GetCommunityDetailById err", zap.Error(err))
			return
		}
		postDetail.CommunityDetail = communityDetail

		postDetailList = append(postDetailList, postDetail)
	}
	return
}
