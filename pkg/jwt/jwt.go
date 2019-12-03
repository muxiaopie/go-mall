package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/muxiaopie/go-mall/pkg/errno"
	"github.com/spf13/viper"
	"time"
)

type Claims struct {
	Id uint `json:"uid"`
	jwt.StandardClaims
}

// 生成TOKEN
func GenerateToken(uid uint) (string, error) {
	//nowTime := time.Now()
	//expireTime := nowTime.Add(3 * time.Hour)

	nowTime := time.Now().Unix()
	expireTime := nowTime + viper.GetInt64("expireTime")
	secret := viper.GetString("secret")

	claims := Claims{
		uid,
		jwt.StandardClaims{
			//ExpiresAt : expireTime.Unix(),
			ExpiresAt: expireTime,
			Issuer:    viper.GetString("issuer"),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))

	return token, err
}

// 解析Token
func ParseToken(tokenString string) (*Claims, error) {
	secret := viper.GetString("secret")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errno.JwtError("That's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errno.JwtError("Token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errno.JwtError("Token not active yet")
			} else {
				return nil, errno.JwtError("Couldn't handle this token:")
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errno.JwtError("Couldn't handle this token:")
}

// 刷新token
func RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	secret := viper.GetString("secret")
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return GenerateToken(claims.Id)
	}
	return "", errno.JwtError("Couldn't handle this token:")
}
