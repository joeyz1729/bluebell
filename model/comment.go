package model

import "time"

type Comment struct {
	CommentId       uint64 `json:"comment_id,string" db:"comment_id"`
	AuthorId uint64 `json:"author_id" db:"author_id"`
	PostId uint64 `json:"post_id" db:"post_id" binding:"required"`
	Title       string `json:"title" db:"title" binding:"required"`
	

	Status     int32     `json:"status" db:"status"`
	CreateTime time.Time `json:"-" db:"create_time"`
	UpdateTime time.Time `json:"-" db:"update_time"`
	
}

type CommentDetail struct {
	CommunityId   uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
	Content     string `json:"content" db:"content" binding:"required"`
	Cancel        uint8
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}
