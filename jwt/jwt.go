package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// 生成token
func GenToken(expire time.Duration, issuer string, jwtSecret string) (string, error) {
	expireTime := time.Now().Add(expire)
	standardClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    issuer,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, standardClaims)
	return tokenClaims.SignedString([]byte(jwtSecret))
	//return tokenClaims.SignedString(r.JwtSecret)
}

// 解析token
func ParseToken(token string, jwtSecret string) (*jwt.StandardClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("wrong type")
	}

	if !tokenClaims.Valid {
		return claims, errors.New("token valid is false")
	}

	//
	return claims, nil
}
