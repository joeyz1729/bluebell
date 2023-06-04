package model

type User struct {
	UserId   uint64 `json:"user_id,string" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`

	AccessToken  string
	RefreshToken string
}

func CreateUser(userID uint64, sf *SignupForm) *User {
	return new(User)
}
