package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/YiZou89/bluebell/model"
)

// Login 验证用户登陆信息
func Login(user *model.User) (err error) {
	inputPassword := user.Password

	sqlStr := `select user_id, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrorUserNotExist
		}
		return ErrorQueryUserData
	}

	if encryptPassword(inputPassword) != user.Password {
		return ErrorIncorrectPassword
	}

	return
}

// CheckUserExist 当用户登陆时，检查用户名是否已存在
func CheckUserExist(sf *model.SignupForm) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err = db.Get(&count, sqlStr, sf.Username)
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 添加存储用户信息
func InsertUser(user *model.User) (err error) {
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	encryptedPassword := encryptPassword(user.Password)
	_, err = db.Exec(sqlStr, user.UserId, user.Username, encryptedPassword)
	return

}

// encryptPassword 将用户密码加密
func encryptPassword(originPassword string) (encryptedPassword string) {
	h := md5.New()
	h.Write([]byte(secret))
	encryptedPassword = hex.EncodeToString(h.Sum([]byte(originPassword)))
	return
}

// GetUserById 通过用户id获取详细信息
func GetUserById(uid uint64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `select username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	if err != nil {
		return
	}
	return
}
