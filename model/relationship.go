package model

type Join struct {
	//Id          uint64 `json:"comment_id,string" db:"comment_id"`
	CommunityId uint64 `json:"community_id,string" db:"community_id"`
	UserId      uint64 `json:"user_id,string" db:"user_id"`
	Cancel      bool   `json:"cancel,int"`
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
