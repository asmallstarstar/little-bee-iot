package jwt

import (
	"time"

	yaml "util/yaml/service"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int32 `json:"userId"`
	jwt.StandardClaims
}

func GenerateToken(userId int32) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(yaml.Yaml.Service.TokenExpireTime) * time.Second)

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(yaml.Yaml.Service.JwtSecret))

	return token, err
}

// return user id in token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(yaml.Yaml.Service.JwtSecret), nil
	})

	if tokenClaims != nil && err == nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
