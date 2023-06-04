package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"zouyi/bluebell/model"
)

func Login(user *model.User) (err error) {
	// 1. get user info from mysql database
	sqlStr := `select user_id,password from user where username = ?`
	inputPassword := user.Password
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

// CheckUserExist before store user's signup information, check if user already exists in mysql database by model.SignupForm
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

// InsertUser insert into mysql database by model.User struct.
func InsertUser(user *model.User) (err error) {
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	encryptedPassword := encryptPassword(user.Password)
	_, err = db.Exec(sqlStr, user.UserId, user.Username, encryptedPassword)
	return

}

// encryptPassword encrypt user's password by md5 and secret salt before insert into database.
func encryptPassword(originPassword string) (encryptedPassword string) {
	h := md5.New()
	h.Write([]byte(secret))
	encryptedPassword = hex.EncodeToString(h.Sum([]byte(originPassword)))
	return
}

// GetUserById select user info (username, email, age etc.)
func GetUserById(uid uint64) (user *model.User, err error) {
	user = new(model.User)
	sqlStr := `select username from user where user_id = ?`
	err = db.Get(user, sqlStr, uid)
	if err != nil {
		return
	}
	return
}
