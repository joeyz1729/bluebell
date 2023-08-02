package logic

import (
	"context"

	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/model"
)

type UserService interface {
	Info(ctx context.Context, uid uint64) (string, error)
}

type userService struct {
}

func (us userService) Info(ctx context.Context, uid uint64) (username string, err error) {
	var user = new(model.User)
	user, err = mysql.GetUserById(uid)
	if err != nil {
		return "unknown user", err
	}
	return user.Username, nil

}

func NewUserService() UserService {
	return &userService{}
}
