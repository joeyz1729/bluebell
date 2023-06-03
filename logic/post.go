package logic

import (
	"zouyi/bluebell/dao/mysql"
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
	// 2. store into database
	if err = mysql.CreatePost(p); err != nil {
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

//localhost:8081/api/v1/post/463364212119306241

func GetPostList(page, size uint64) (postDetailList []*model.PostDetail, err error) {
	postDetailList = make([]*model.PostDetail, 0, 2)
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
