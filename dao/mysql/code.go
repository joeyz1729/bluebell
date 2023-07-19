package mysql

import "errors"

const secret = "bluebell.joey1729"

var (
	ErrorUserExist         = errors.New("user already exists")
	ErrorUserNotExist      = errors.New("user not exist")
	ErrorQueryUserData     = errors.New("query user data err")
	ErrorIncorrectPassword = errors.New("incorrect password")

	ErrInvalidId   = errors.New("invalid id")
	ErrQueryFailed = errors.New("select query failed")

	ErrInsertPost = errors.New("insert post info failed")
)
