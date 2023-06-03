package model

import (
	"time"
)

type Post struct {
	Id       uint64 `json:"post_id,string" db:"post_id"`
	AuthorId uint64 `json:"author_id" db:"author_id"`

	CommunityId uint64 `json:"community_id" db:"community_id" binding:"required"`
	Title       string `json:"title" db:"title" binding:"required"`
	Content     string `json:"content" db:"content" binding:"required"`

	Status     int32     `json:"status" db:"status"`
	CreateTime time.Time `json:"-" db:"create_time"`
	UpdateTime time.Time `json:"-" db:"update_time"`
}

type PostDetail struct {
	AuthorName string `json:"username" db:"username"`
	*Post
	*CommunityDetail
}
