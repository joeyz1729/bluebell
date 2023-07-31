package model

type Community struct {
	CommunityId   uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
}

type CommunityDetail struct {
	CommunityId   uint64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
	Introduction  string `json:"introduction,omitempty" db:"introduction"`

	JoinCount int64
	PostCount int64
	IsJoined  bool
	//CreateTime    time.Time `json:"create_time" db:"create_time"`
}
