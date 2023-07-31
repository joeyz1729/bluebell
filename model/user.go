package model

type User struct {
	UserId   uint64 `json:"user_id,string" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`

	// email, phone, avatar, background_image, ...

	AccessToken  string
	RefreshToken string
}

type UserDetail struct {
	Id   uint64 `json:"user_id,string" db:"user_id"`
	Name string `json:"username" db:"username"`

	FollowCount   int64
	FollowerCount int64
	IsFollow      bool

	// 发帖数
	WorkCount int64

	// 帖子被赞数
	TotalFavorited int64 // auth id - post id - vote

	// 加入的社区数量
	JoinedCount int64
}

func CreateUser(userID uint64, sf *SignupForm) *User {
	return new(User)
}
