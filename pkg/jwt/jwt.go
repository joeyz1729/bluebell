package jwt

import (
	"fmt"
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
	AccessTokenExpireDuration  = 30 * 24 * time.Hour
	RefreshTokenExpireDuration = time.Hour * 24 * 30
)

var (
	ErrorInvalidToken = errors.New("invalid token")
)

// salt
var secret = []byte("joey89")

func GenToken(userId uint64, username string) (aToken, rToken string, err error) {
	c := &Claims{
		userId,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
	}
	fmt.Println(c.ExpiresAt)
	// with userid and username
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(secret)

	// without any data
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
		Issuer:    "bluebell",
	}).SignedString(secret)

	return
}

func ParseToken(tokenString string) (claims *Claims, err error) {

	claims = new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	fmt.Println(claims.UserId)
	fmt.Println(claims.Username)
	fmt.Println(claims.ExpiresAt)
	fmt.Println(time.Now().Unix())

	if err != nil {
		return nil, err
	}
	if token.Valid {
		return claims, nil
	}
	return nil, ErrorInvalidToken
}

func RefreshToken(accessToken, refreshToken string) (newAccessToken, newRefreshToken string, err error) {
	// invalid token
	if _, err = jwt.Parse(refreshToken, keyFunc); err != nil {
		return
	}

	// parse token
	var claims Claims
	_, err = jwt.ParseWithClaims(accessToken, &claims, keyFunc)

	v, _ := err.(*jwt.ValidationError)

	// token expired
	if v.Errors == jwt.ValidationErrorExpired {
		return GenToken(claims.UserId, claims.Username)
	}
	fmt.Println(3)
	return
}

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return secret, nil
}
