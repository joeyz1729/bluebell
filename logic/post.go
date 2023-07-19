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
	err = redis.CreatePost(p.Id, p.CommunityId)
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
	//3. get community detail
	var communityDetail = new(model.CommunityDetail)
	communityDetail, err = mysql.GetCommunityDetailById(pd.Post.CommunityId)

	if err != nil {
		zap.L().Error("GetCommunityDetailById() failed", zap.Error(err))
		return
	}
	pd.CommunityDetail = communityDetail

	return
}

func GetPostList(page, size int64) (postDetailList []*model.PostDetail, err error) {
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

// GetCommunityPostList 按照执行社区id获取帖子信息
func GetCommunityPostList(cid uint64, page, size int64) (postDetailList []*model.PostDetail, err error) {
	// 获取社区信息
	communityDetail, err := mysql.GetCommunityDetailById(cid)
	if err != nil {
		zap.L().Error("get community detail err", zap.Error(err))
		return nil, err
	}

	// 获取帖子列表
	var posts []*model.Post
	posts, err = mysql.GetCommunityPostList(cid, page, size)
	if err != nil {
		zap.L().Error("get community post list err", zap.Error(err))
		return nil, err
	}

	// 补充帖子详细信息
	for _, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post

		var user = new(model.User)
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			zap.L().Error("get username by id err", zap.Error(err))
			continue
		}
		postDetail.AuthorName = user.Username
		postDetail.CommunityDetail = communityDetail

		postDetailList = append(postDetailList, postDetail)
	}

	return
}

func GetPostListInOrder(form *model.PostsForm) (postDetailList []*model.PostDetail, err error) {
	postDetailList = make([]*model.PostDetail, 0, form.Size)
	var posts []*model.Post
	// 从redis中按照指定顺序获取post id
	ids, err := redis.GetPostIdsInOrder(form)
	if err != nil || len(ids) == 0 {
		zap.L().Error("get post ids in order from redis err, or len(ids) == 0", zap.Error(err))
		return
	}

	// 从数据库中获取指定id顺序的帖子
	posts, err = mysql.GetPostByIds(ids)
	if err != nil {
		zap.L().Error("get post by ids from mysql err", zap.Error(err))
		return nil, err
	}

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	// query post detail
	for id, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post
		postDetail.Votes = voteData[id]

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

func FullPostDetail(posts []*model.Post) (postDetailList []*model.PostDetail, err error) {
	postDetailList = make([]*model.PostDetail, 0, len(posts))
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
