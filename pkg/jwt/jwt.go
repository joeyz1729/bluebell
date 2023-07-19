package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"errors"
)

type Claims struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	AccessTokenExpireDuration  = 30 * 24 * time.Hour // 20 * time.Minute
	RefreshTokenExpireDuration = 24 * 30 * time.Hour // 30 * 24 * time.Hour
)

var (
	ErrorInvalidToken = errors.New("invalid token")
)

// salt
var secret = []byte("joey1729.bluebell")

// GenToken 根据用户身份信息生成jwt 的access token和refresh token
func GenToken(userId uint64, username string) (aToken, rToken string, err error) {
	c := &Claims{
		// 自定义字段
		userId,
		username,
		jwt.StandardClaims{
			// 标准字段
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "bluebell",                                       // 发布者
		},
	}

	// 使用指定加密算法，将json对象c加密生成token，再对token进行签名
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(secret)

	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
		Issuer:    "bluebell",
	}).SignedString(secret)

	return
}

// ParseToken 将token字符串解析为json对象
func ParseToken(tokenString string) (claims *Claims, err error) {
	claims = new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}
	return nil, ErrorInvalidToken
}

// RefreshToken 验证并解析access token和refresh token，如果合法则重新生成
func RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// 解析refresh token并验证是否有效
	if _, err = jwt.Parse(refreshToken, keyFunc); err != nil {
		return
	}

	// 解析access token获取json信息
	var claims Claims
	_, err = jwt.ParseWithClaims(accessToken, &claims, keyFunc)
	v, _ := err.(*jwt.ValidationError)

	// 如果access token 解析成功并且过期
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId, claims.Username)
	}
	return accessToken, refreshToken, nil
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return secret, nil
}
