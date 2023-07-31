package model

type User struct {
	UserId   uint64 `json:"user_id,string" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`

	AccessToken  string
	RefreshToken string
}

type UserDetail struct {
	Id   uint64 `json:"user_id,string" db:"user_id"`
	Name string `json:"username" db:"username"`

	// 关注
	FollowCount   uint64
	FollowerCount uint64
	IsFollow      bool

	// 发帖数
	WorkCount uint64

	// 点赞
	TotalFavorited uint64 // auth id - post id - vote
	FavoriteCount  uint64

	//Avatar          string
	//BackgroundImage string
	//Signature       string
}

func CreateUser(userID uint64, sf *SignupForm) *User {
	return new(User)
}
