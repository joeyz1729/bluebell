package model

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

type SignupForm struct {
	Age             uint8  `json:"age" binding:"omitempty,gte=1,lte=130"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Email           string `json:"email,omitempty" binding:"omitempty,email"`
}

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RefreshForm struct {
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Email           string `json:"email,omitempty" binding:"omitempty,email"`
}

type VoteForm struct {
	// UserID, get from token
	PostID   string `json:"post_id" binding:"required"`
	Attitude int8   `json:"attitude,string" binding:"oneof=1 0 -1" `
}

type PostsForm struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Order string `form:"order"`
}

type CommunityPostsForm struct {
	CommunityId uint64 `json:"community_id" form:"community_id"`
	*PostsForm
}
