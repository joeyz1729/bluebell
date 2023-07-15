package logic

import (
	"zouyi/bluebell/dao/mysql"
	"zouyi/bluebell/model"
	"zouyi/bluebell/pkg/jwt"
	"zouyi/bluebell/pkg/snowflake"
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
	// 4. insert into database
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

	// 生成jwt token
	accessToken, refreshToken, err := jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return nil, err
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}
