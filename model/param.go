package model

type SignupForm struct {
	//Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	//Email      string `json:"email" binding:"required,email"`
}

type LoginForm struct {
	//Age        uint8  `json:"age" binding:"gte=1,lte=130"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	//Email      string `json:"email" binding:"required,email"`
}

type VoteForm struct {
	// UserID, get from token
	PostID   string `json:"post_id" binding:"required"`
	Attitude int8   `json:"attitude,string" binding:"oneof=1 0 -1" ` // 赞成票(1)还是反对票(-1)取消投票(0)
}
