package logic

import (
	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/dao/redis"
	"github.com/YiZou89/bluebell/model"
	"github.com/YiZou89/bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

// CreatePost 创建帖子信息并存储
func CreatePost(p *model.Post) (err error) {
	// 生成帖子id
	id, err := snowflake.GenID()
	if err != nil {
		return
	}
	p.Id = id

	// 创建并保存
	err = redis.CreatePost(p.Id, p.CommunityId)
	if err != nil {
		zap.L().Error("post info save to redis err", zap.Error(err))
		return
	}
	err = mysql.CreatePost(p)
	if err != nil {
		zap.L().Error("post info save to mysql err", zap.Error(err))
		return
	}
	return
}

// GetPostDetailById 获取指定帖子信息
func GetPostDetailById(pid uint64) (pd *model.PostDetail, err error) {
	pd = new(model.PostDetail)

	// 获取帖子信息
	var post = new(model.Post)
	post, err = mysql.GetPostById(pid)
	if err != nil {
		return
	}
	pd.Post = post

	// 获取发帖人信息
	user, err := mysql.GetUserById(pd.AuthorId)
	if err != nil {
		return
	}
	pd.AuthorName = user.Username

	// 获取社区信息
	var communityDetail = new(model.CommunityDetail)
	communityDetail, err = mysql.GetCommunityDetailById(pd.Post.CommunityId)

	if err != nil {
		return
	}
	pd.CommunityDetail = communityDetail

	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (postDetailList []*model.PostDetail, err error) {
	var posts []*model.Post
	posts, err = mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}

	for _, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post

		// 获取用户名称
		var user = new(model.User)
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			return
		}
		postDetail.AuthorName = user.Username

		// 获取社区详细信息
		var communityDetail = new(model.CommunityDetail)
		communityDetail, err = mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
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
		return nil, err
	}

	// 获取帖子列表
	var posts []*model.Post
	posts, err = mysql.GetCommunityPostList(cid, page, size)
	if err != nil {
		return nil, err
	}

	// 补充帖子详细信息
	for _, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post

		var user = new(model.User)
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			continue
		}
		postDetail.AuthorName = user.Username
		postDetail.CommunityDetail = communityDetail

		postDetailList = append(postDetailList, postDetail)
	}

	return
}

func GetPublishPostList(uid uint64, page, size int64) (postDetailList []*model.PostDetail, err error) {
	// 获取帖子列表
	var posts []*model.Post
	posts, err = mysql.GetUserPostList(uid, page, size)
	if err != nil {
		return nil, err
	}

	// 补充帖子详细信息
	for _, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post

		var user = new(model.User)
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			continue
		}

		postDetail.AuthorName = user.Username

		postDetailList = append(postDetailList, postDetail)
	}

	return
	// 获取社区信息

}

func GetFavoritePostList(uid uint64, page, size int64) (postDetailList []*model.PostDetail, err error) {
	// 获取帖子列表
	//TODO
	var posts []*model.Post
	posts, err = mysql.GetUserPostList(uid, page, size)
	if err != nil {
		return nil, err
	}

	// 补充帖子详细信息
	for _, post := range posts {
		var postDetail = new(model.PostDetail)
		postDetail.Post = post

		var user = new(model.User)
		user, err = mysql.GetUserById(post.AuthorId)
		if err != nil {
			continue
		}

		postDetail.AuthorName = user.Username

		postDetailList = append(postDetailList, postDetail)
	}

	return
	// 获取社区信息

}

// GetPostListInOrder 按照指定顺序获取帖子列表
func GetPostListInOrder(order string, page, size int64) (postDetailList []*model.PostDetail, err error) {
	postDetailList = make([]*model.PostDetail, 0, size)
	var posts []*model.Post

	// 从redis中按照指定顺序获取post id
	ids, err := redis.GetPostIdsInOrder(order, page, size)
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
