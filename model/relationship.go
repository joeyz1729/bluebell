package model

type Join struct {
	Id          uint64
	CommunityId uint64
	UserId      uint64
}

type Follow struct {
	Id         uint64
	UserId     uint64
	FollowerId uint64
	Cancel     bool
}

type Vote struct {
	Id     uint64
	PostId uint64
	UserId uint64
	Cancel bool
}
