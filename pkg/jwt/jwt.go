package jwt

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"

	"errors"
)

type Claims struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

var (
	AccessTokenExpireDuration  = viper.GetDuration("jwt_expire") * time.Hour
	RefreshTokenExpireDuration = time.Hour * 25 * 7
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
	var token *jwt.Token
	claims = new(Claims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
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
	fmt.Println(1)
	// parse token
	var claims Claims
	_, err = jwt.ParseWithClaims(accessToken, &claims, keyFunc)
	fmt.Println(2)
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
