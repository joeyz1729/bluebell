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

	// 关注
	//FollowCount   uint64
	//FollowerCount uint64
	//IsFollow      bool

	// 发帖数
	WorkCount uint64

	// 帖子被赞数
	TotalFavorited uint64 // auth id - post id - vote

	// 加入的社区数量
	JoinedCount uint64
}

func CreateUser(userID uint64, sf *SignupForm) *User {
	return new(User)
}
