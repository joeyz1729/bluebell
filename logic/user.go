package logic

import (
	"strconv"

	"github.com/YiZou89/bluebell/dao/mysql"
	"github.com/YiZou89/bluebell/dao/redis"
	"github.com/YiZou89/bluebell/model"
	"go.uber.org/zap"

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

func GetUserDetailById(uid, userId uint64) (user *model.UserDetail, err error) {
	user = new(model.UserDetail)
	user, err = redis.GetUserInfo(strconv.Itoa(int(uid)), strconv.Itoa(int(userId)))
	if err == nil {
		return user, err
	}
	zap.L().Error("get user info from redis err", zap.Error(err))

	user, err = mysql.GetUserDetailById(uid, userId)
	if err != nil {
		zap.L().Error("get user info from mysql err", zap.Error(err))
		return nil, err
	}
	return user, nil
}
