package controller

type CodeType int64

const (
	CodeSuccess CodeType = 1000 + iota
	CodeInvalidParams
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeInvalidToken
	CodeInvalidAuthFormat
	CodeNotLogin
	ErrVoteRepeated
	ErrorVoteTimeExpire
)

var msgFlags = map[CodeType]string{
	CodeSuccess:         "success",
	CodeInvalidParams:   "invalid params",
	CodeUserExist:       "username already exists",
	CodeUserNotExist:    "user not exists",
	CodeInvalidPassword: "invalid password",
	CodeServerBusy:      "server busy",

	CodeInvalidToken:      "invalid token",
	CodeInvalidAuthFormat: "invalid auth format",
	CodeNotLogin:          "not login",

	ErrVoteRepeated:     "请勿重复投票",
	ErrorVoteTimeExpire: "投票时间已过",
}

func (code CodeType) Msg() string {
	msg, ok := msgFlags[code]
	if !ok {
		msg = msgFlags[CodeServerBusy]
	}
	return msg

}
