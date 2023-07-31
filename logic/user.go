package logic

import (
	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/model"

	"github.com/YiZou89/bluebell/pkg/jwt"
	"github.com/YiZou89/bluebell/pkg/snowflake"
)

func Signup(sf *model.SignupForm) (err error) {
	err = mysql.CheckUserExist(sf)
	if err != nil {
		return err
	}

	userID, err := snowflake.GenID()
	if err != nil {
		// TODO
		// const ErrorGenIdFailed
		return err
	}

	user := &model.User{
		UserId:   userID,
		Username: sf.Username,
		Password: sf.Password,
	}

	return mysql.InsertUser(user)
}

func Login(lf *model.LoginForm) (user *model.User, err error) {
	user = &model.User{
		Username: lf.Username,
		Password: lf.Password,
	}
	if err = mysql.Login(user); err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return nil, err
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}

func GetUserDetailById(uid uint64) (userDetail model.UserDetail, err error) {
	userDetail = model.UserDetail{}
	userDetail.Id = uid

	user, err := mysql.GetUserById(uid)
	if err != nil {
		return
	}
	userDetail.Name = user.Username

	return userDetail, nil
}
